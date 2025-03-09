package model

import (
	"context"
	"messenger/internal/mto"
)

type repo interface{
	CreateChat(context.Context, mto.Chat) (string, error)
	CreateMessage(context.Context, mto.Message) (error)
	GetMessagesFromChat(context.Context, string) ([]mto.Message, error)
}

type Model struct{
	repo repo
}

func NewModel(repo repo) *Model {
	return &Model{repo:repo}
}

func (m Model) InitSession(context.Context) (string, error){
	return "",nil
}

func (m Model) CreateChat(ctx context.Context, chat mto.Chat) (string, error){
	m.repo.CreateChat(ctx,chat)
	return "",nil
}

func (m Model) SendMessageToChat(ctx context.Context, msg mto.Message) (error){
	m.repo.CreateMessage(ctx,msg)
	return nil
}

func (m Model) GetMessagesFromChat(ctx context.Context,id string) ([]mto.Message, error) {
	m.repo.GetMessagesFromChat(ctx,id)
	return nil,nil
}