package main

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/rubberduckkk/go-demo/proto/demo"
)

type DemoService struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewDemoService(ctx context.Context) *DemoService {
	ctx, cancel := context.WithCancel(ctx)
	return &DemoService{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (d *DemoService) ClientStream(server demo.Demo_ClientStreamServer) error {
	d.wg.Add(1)
	defer d.wg.Done()

	logrus.Info("client stream accessed")
	defer func() {
		logrus.Info("client stream closed")
	}()

	var cnt int64
	for {
		select {
		case <-d.ctx.Done():
			logrus.Info("demo service shutdown")
			return server.SendAndClose(&demo.Pong{Ack: cnt, Timestamp: time.Now().UnixMilli()})
		default:
			in, err := server.Recv()
			if err == io.EOF {
				return server.SendAndClose(&demo.Pong{Ack: cnt, Timestamp: time.Now().UnixMilli()})
			}
			if err != nil {
				logrus.WithError(err).Error("demo server receive error")
				return err
			}

			logrus.WithField("in", in).Info("demo client receive")
			cnt++
		}
	}
}

func (d *DemoService) shutdown() {
	d.cancel()
	d.wg.Wait()
}
