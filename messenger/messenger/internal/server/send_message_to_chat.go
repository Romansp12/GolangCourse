package server

import (
	"context"
	"fmt"

	"messenger/internal/mto"
	desc "messenger/pkg/api/messenger/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Server) SendMessageToChat(ctx context.Context,in *desc.SendMessageToChatReq) (*emptypb.Empty, error) {
	fmt.Println("SendMessageToChat", in.Id)
	s.model.SendMessageToChat(ctx,mto.Message{})
	return &emptypb.Empty{},nil
}