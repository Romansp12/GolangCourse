package server

import (
	"context"
	"fmt"

	"messenger/internal/mto"
	desc "messenger/pkg/api/messenger/v1"
)

func (s Server) CreateChat(ctx context.Context,in *desc.CreateChatReq) (*desc.CreateChatResp, error) {
	fmt.Println("CreateChat user ",in.Id)
	s.model.CreateChat(ctx,mto.Chat{})
	return &desc.CreateChatResp{Id: "456"},nil
}