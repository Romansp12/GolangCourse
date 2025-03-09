package server

import (
	"context"
	"fmt"

	desc "messenger/pkg/api/messenger/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Server) InitSession(ctx context.Context,in *emptypb.Empty) (*desc.InitSessionResp, error) {
	fmt.Println("InitSession")
	s.model.InitSession(ctx)
	return &desc.InitSessionResp{Id: "123"},nil
}