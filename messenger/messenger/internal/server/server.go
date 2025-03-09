package server

import (
	"context"
	desc "messenger/pkg/api/messenger/v1"
	"messenger/internal/mto"
)

type model interface{
	InitSession(context.Context) (string, error)
	CreateChat(context.Context, mto.Chat) (string, error)
	SendMessageToChat(context.Context, mto.Message) (error)
	GetMessagesFromChat(context.Context, string) ([]mto.Message, error)
}

type Server struct {
	desc.UnimplementedMESSENGERServer
	model model
}

func NewServer(model model) *Server{
	return &Server{
		model: model,
	}
}


