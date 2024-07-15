package main

import (
	"context"
	"math/rand"
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/rubberduckkk/go-demo/grpc/metadata/pb"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

func BenchmarkRequestComposite(b *testing.B) {
	cc, err := grpc.DialInsecure(context.Background(), option.WithEndpoint("localhost:8789"))
	if err != nil {
		b.Fatal(err)
	}
	defer cc.Close()
	head := &pb.RequestHead{
		IsInternal: false,
		ServerType: 1,
		Route:      "x",
		RequestId:  "y",
		TraceId:    "z",
		Session:    randSeq(50),
		Extra:      randSeq(50),
	}
	body := &pb.RequestComposite{
		Head: head,
		Body: randSeq(500),
	}
	cli := pb.NewDemoServiceClient(cc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cli.DoRequestComposite(context.Background(), body)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRequest(b *testing.B) {
	cc, err := grpc.DialInsecure(context.Background(), option.WithEndpoint("localhost:8789"))
	if err != nil {
		b.Fatal(err)
	}
	defer cc.Close()
	head := &pb.RequestHead{
		IsInternal: false,
		ServerType: 1,
		Route:      "x",
		RequestId:  "y",
		TraceId:    "z",
		Session:    randSeq(50),
		Extra:      randSeq(50),
	}

	body := &pb.Request{
		Body: randSeq(500),
	}

	m := metadata.New(map[string]string{
		"head": proto.MarshalTextString(head),
	})
	ctx := metadata.NewIncomingContext(context.Background(), m)
	cli := pb.NewDemoServiceClient(cc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cli.DoRequest(ctx, body)
		if err != nil {
			b.Fatal(err)
		}
	}
}
