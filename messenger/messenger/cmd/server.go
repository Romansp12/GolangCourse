package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "messenger/pkg/api/messenger/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	fmt.Println("start server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	cur:=tmp{}
	desc.RegisterMESSENGERServer(grpcServer, cur)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			fmt.Println(err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}

type tmp struct {
}

func (t tmp) InitSession(context.Context, *emptypb.Empty) (*desc.InitSessionResp, error) {
	return &desc.InitSessionResp{Id:"123"},nil
}

func (t *tmp) mustEmbedUnimplementedMESSENGERServer() {
}
