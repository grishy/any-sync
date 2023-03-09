package ocache

import (
	"context"
	"errors"
	"github.com/anytypeio/any-sync/app/logger"
	"github.com/anytypeio/any-sync/util/slice"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	ErrClosed    = errors.New("object cache closed")
	ErrExists    = errors.New("object exists")
	ErrNotExists = errors.New("object not exists")
)

var (
	defaultTTL = time.Minute
	defaultGC  = 20 * time.Second
)

var log = logger.NewNamed("ocache")

type LoadFunc func(ctx context.Context, id string) (value Object, err error)

type Option func(*oCache)

var WithLogger = func(l *zap.SugaredLogger) Option {
	return func(cache *oCache) {
		cache.log = l
	}
}

var WithTTL = func(ttl time.Duration) Option {
	return func(cache *oCache) {
		cache.ttl = ttl
	}
}

var WithGCPeriod = func(gcPeriod time.Duration) Option {
	return func(cache *oCache) {
		cache.gc = gcPeriod
	}
}

func New(loadFunc LoadFunc, opts ...Option) OCache {
	c := &oCache{
		data:     make(map[string]*entry),
		loadFunc: loadFunc,
		timeNow:  time.Now,
		ttl:      defaultTTL,
		gc:       defaultGC,
		closeCh:  make(chan struct{}),
		log:      log.Sugar(),
	}
	for _, o := range opts {
		if o != nil {
			o(c)
		}
	}
	if c.ttl != 0 && c.gc != 0 {
		go c.ticker()
	}
	return c
}

type Object interface {
	Close() (err error)
	TryClose() (res bool, err error)
}

type OCache interface {
	// DoLockedIfNotExists does an action if the object with id is not in cache
	// under a global lock, this will prevent a race which otherwise occurs
	// when object is created in parallel with action
	DoLockedIfNotExists(id string, action func() error) error
	// Get gets an object from cache or create a new one via 'loadFunc'
	// Increases the object refs counter on successful
	// When 'loadFunc' returns a non-nil error, an object will not be stored to cache
	Get(ctx context.Context, id string) (value Object, err error)
	// Pick returns value if it's presents in cache (will not call loadFunc)
	Pick(ctx context.Context, id string) (value Object, err error)
	// Add adds new object to cache
	// Returns error when object exists
	Add(id string, value Object) (err error)
	// Remove closes and removes object
	Remove(id string) (ok bool, err error)
	// ForEach iterates over all loaded objects, breaks when callback returns false
	ForEach(f func(v Object) (isContinue bool))
	// GC frees not used and expired objects
	// Will automatically called every 'gcPeriod'
	GC()
	// Len returns current cache size
	Len() int
	// Close closes all objects and cache
	Close() (err error)
}

type oCache struct {
	mu       sync.Mutex
	data     map[string]*entry
	loadFunc LoadFunc
	timeNow  func() time.Time
	ttl      time.Duration
	gc       time.Duration
	closed   bool
	closeCh  chan struct{}
	log      *zap.SugaredLogger
	metrics  *metrics
}

func (c *oCache) Get(ctx context.Context, id string) (value Object, err error) {
	var (
		e    *entry
		ok   bool
		load bool
	)
Load:
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil, ErrClosed
	}
	if e, ok = c.data[id]; !ok {
		e = newEntry(id, nil, entryStateLoading)
		load = true
		c.data[id] = e
	}
	c.mu.Unlock()
	reload, err := e.waitClose(ctx, id)
	if err != nil {
		return nil, err
	}
	if reload {
		goto Load
	}
	if load {
		go c.load(ctx, id, e)
	}
	c.metricsGet(!load)
	return e.waitLoad(ctx, id)
}

func (c *oCache) metricsGet(hit bool) {
	if c.metrics == nil {
		return
	}
	if hit {
		c.metrics.hit.Inc()
	} else {
		c.metrics.miss.Inc()
	}
}

func (c *oCache) Pick(ctx context.Context, id string) (value Object, err error) {
	c.mu.Lock()
	val, ok := c.data[id]
	if !ok || val.isClosing() {
		c.mu.Unlock()
		return nil, ErrNotExists
	}
	c.mu.Unlock()
	c.metricsGet(true)
	return val.waitLoad(ctx, id)
}

func (c *oCache) load(ctx context.Context, id string, e *entry) {
	defer close(e.load)
	value, err := c.loadFunc(ctx, id)

	c.mu.Lock()
	defer c.mu.Unlock()
	if err != nil {
		e.loadErr = err
		delete(c.data, id)
	} else {
		e.value = value
		e.setActive(false)
	}
}

func (c *oCache) Remove(id string) (ok bool, err error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		err = ErrClosed
		return
	}
	e, ok := c.data[id]
	if !ok {
		c.mu.Unlock()
		return
	}
	c.mu.Unlock()
	return c.remove(e, true)
}

func (c *oCache) remove(e *entry, remData bool) (ok bool, err error) {
	<-e.load
	if e.value == nil {
		return false, ErrNotExists
	}
	_, curState := e.setClosing(true)
	if curState == entryStateClosing {
		err = e.value.Close()
		e.setClosed()
	}
	if !remData {
		return
	}
	c.mu.Lock()
	if curState == entryStateClosing {
		delete(c.data, e.id)
	}
	c.mu.Unlock()
	return
}

func (c *oCache) DoLockedIfNotExists(id string, action func() error) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return ErrClosed
	}
	if _, ok := c.data[id]; ok {
		return ErrExists
	}
	return action()
}

func (c *oCache) Add(id string, value Object) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[id]; ok {
		return ErrExists
	}
	e := newEntry(id, value, entryStateActive)
	close(e.load)
	c.data[id] = e
	return
}

func (c *oCache) ForEach(f func(obj Object) (isContinue bool)) {
	var objects []Object
	c.mu.Lock()
	for _, v := range c.data {
		select {
		case <-v.load:
			if v.value != nil && !v.isClosing() {
				objects = append(objects, v.value)
			}
		default:
		}
	}
	c.mu.Unlock()
	for _, obj := range objects {
		if !f(obj) {
			return
		}
	}
}

func (c *oCache) ticker() {
	ticker := time.NewTicker(c.gc)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.GC()
		case <-c.closeCh:
			return
		}
	}
}

func (c *oCache) GC() {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return
	}
	deadline := c.timeNow().Add(-c.ttl)
	var toClose []*entry
	for _, e := range c.data {
		if e.getState() != entryStateActive {
			continue
		}
		if e.lastUsage.Before(deadline) {
			e.close = make(chan struct{})
			toClose = append(toClose, e)
		}
	}
	size := len(c.data)
	c.mu.Unlock()

	for idx, e := range toClose {
		prevState, _ := e.setClosing(false)
		if prevState == entryStateClosing || prevState == entryStateClosed {
			toClose[idx] = nil
			continue
		}
		ok, err := e.value.TryClose()
		if !ok {
			e.setActive(true)
			toClose[idx] = nil
			continue
		} else {
			e.setClosed()
		}
		if err != nil {
			c.log.With("object_id", e.id).Warnf("GC: object close error: %v", err)
		}
	}
	toClose = slice.DiscardFromSlice(toClose, func(e *entry) bool {
		return e == nil
	})
	c.log.Infof("GC: removed %d; cache size: %d", len(toClose), size)
	if len(toClose) > 0 && c.metrics != nil {
		c.metrics.gc.Add(float64(len(toClose)))
	}
	c.mu.Lock()
	for _, e := range toClose {
		delete(c.data, e.id)
	}
	c.mu.Unlock()
}

func (c *oCache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.data)
}

func (c *oCache) Close() (err error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return ErrClosed
	}
	c.closed = true
	close(c.closeCh)
	var toClose []*entry
	for _, e := range c.data {
		toClose = append(toClose, e)
	}
	c.mu.Unlock()
	for _, e := range toClose {
		if _, err := c.remove(e, false); err != ErrNotExists {
			c.log.With("object_id", e.id).Warnf("cache close: object close error: %v", err)
		}
	}
	return nil
}
