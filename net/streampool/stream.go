package streampool

import (
	"context"
	"sync/atomic"

	"github.com/cheggaaa/mb/v3"
	"go.uber.org/zap"
	"storj.io/drpc"

	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/util/multiqueue"
)

type stream struct {
	peerId       string
	peerCtx      context.Context
	stream       drpc.Stream
	pool         *streamPool
	streamId     uint32
	closed       atomic.Bool
	l            logger.CtxLogger
	queue        *multiqueue.Queue[drpc.Message]
	stats        streamStat
	syncDelegate StreamSyncDelegate
}

func (sr *stream) write(msg drpc.Message) (err error) {
	sr.stats.AddMessage(msg)
	err = sr.queue.TryAdd(msg)
	if err != nil {
		sr.stats.RemoveMessage(msg)
	}
	return err
}

func (sr *stream) readLoop() error {
	defer func() {
		sr.streamClose()
	}()
	sr.l.Debug("stream read started")
	for {
		msg := sr.pool.syncDelegate.NewReadMessage()
		if err := sr.stream.MsgRecv(msg, EncodingProto); err != nil {
			sr.l.Info("msg receive error", zap.Error(err))
			return err
		}
		ctx := streamCtx(sr.peerCtx, sr.streamId, sr.peerId)
		ctx = logger.CtxWithFields(ctx, zap.String("peerId", sr.peerId))
		if err := sr.pool.syncDelegate.HandleMessage(ctx, sr.peerId, msg); err != nil {
			sr.l.Info("msg handle error", zap.Error(err))
			return err
		}
	}
}

func (sr *stream) writeLoop() {
	for {
		msg, err := sr.queue.WaitOne(sr.peerCtx)
		if err != nil {
			if err != mb.ErrClosed {
				sr.streamClose()
			}
			return
		}
		if err := sr.stream.MsgSend(msg, EncodingProto); err != nil {
			sr.l.Warn("msg send error", zap.Error(err))
			sr.streamClose()
			return
		}
		sr.stats.RemoveMessage(msg)
	}
}

func (sr *stream) streamClose() {
	if !sr.closed.Swap(true) {
		_ = sr.syncDelegate.RemoveQueue(sr.peerId)
		_ = sr.stream.Close()
		sr.pool.removeStream(sr.streamId)
	}
}
