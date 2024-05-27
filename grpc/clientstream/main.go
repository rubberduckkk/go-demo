package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/rubberduckkk/go-demo/proto/demo"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	ctx := context.Background()
	s := NewDemoService(ctx)
	demo.RegisterDemoServer(g, s)

	handleSignal := func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		for sig := range sigChan {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				logrus.WithField("signal", sig).Info("received signal")
				s.shutdown()
				logrus.Info("demo service shutdown finished")
				g.GracefulStop()
				logrus.Info("graceful shutdown finished")
				return
			default:
				logrus.WithField("signal", sig).Warn("received signal")
			}
		}
	}

	go handleSignal()

	err = g.Serve(ln)
	if err != nil {
		logrus.WithError(err).Error("gRPC server failed")
		return
	}
}
