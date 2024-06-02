package streampool

import (
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/debugstat"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/metric"
)

const CName = "common.net.streampool"

var log = logger.NewNamed(CName)

func New() Service {
	return new(service)
}

type StreamConfig struct {
	// SendQueueSize size of the queue for write per peer
	SendQueueSize int
	// DialQueueWorkers how many workers will dial to peers
	DialQueueWorkers int
	// DialQueueSize size of the dial queue
	DialQueueSize int
}

type Service interface {
	NewStreamPool(h StreamOpener, conf StreamConfig) StreamPool
	app.Component
}

type service struct {
	metric    metric.Metric
	debugStat debugstat.StatService
}

func (s *service) NewStreamPool(h StreamOpener, conf StreamConfig) StreamPool {
	pl := NewExecPool(conf.DialQueueWorkers, conf.DialQueueSize)
	sp := &streamPool{
		streamOpener:    h,
		writeQueueSize:  conf.SendQueueSize,
		streamIdsByPeer: map[string][]uint32{},
		streamIdsByTag:  map[string][]uint32{},
		streams:         map[uint32]*stream{},
		opening:         map[string]*openingProcess{},
		dial:            pl,
		statService:     s.debugStat,
	}
	sp.statService.AddProvider(sp)
	pl.Run()
	if s.metric != nil {
		registerMetrics(s.metric.Registry(), sp, "")
	}
	return sp
}

func (s *service) Init(a *app.App) (err error) {
	s.metric, _ = a.Component(metric.CName).(metric.Metric)
	s.debugStat, _ = a.Component(debugstat.CName).(debugstat.StatService)
	if s.debugStat == nil {
		s.debugStat = debugstat.NewNoOp()
	}
	return nil
}

func (s *service) Name() (name string) {
	return CName
}
