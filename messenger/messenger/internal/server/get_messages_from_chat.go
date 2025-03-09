package server

import (
	"context"
	"fmt"

	desc "messenger/pkg/api/messenger/v1"
)

func (s Server) GetMessagesFromChat(ctx context.Context,in *desc.GetMessagesFromChatReq) (*desc.GetMessagesFromChatResp, error) {
	fmt.Println("GetMessagesFromChat ", in.Id)
	s.model.GetMessagesFromChat(ctx,in.Id)
	return &desc.GetMessagesFromChatResp{},nil
}