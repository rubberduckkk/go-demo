package main

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rubberduckkk/go-demo/grpc/metadata/pb"
)

func newService() *service {
	return &service{}
}

type service struct {
}

func (s *service) DoRequestComposite(ctx context.Context, composite *pb.RequestComposite) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *service) DoRequest(ctx context.Context, request *pb.Request) (*emptypb.Empty, error) {
	return nil, nil
}
