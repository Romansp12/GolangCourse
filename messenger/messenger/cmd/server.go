package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "messenger/pkg/api/messenger/v1"
	"messenger/internal/server"
	"messenger/internal/repository"
	"messenger/internal/model"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	repository:=repository.Repository{}
	model:=model.NewModel(repository)
	server:=server.NewServer(model)
	desc.RegisterMESSENGERServer(grpcServer, server)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			fmt.Println(err)
		}
	}()
	
	fmt.Println("start server")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}

