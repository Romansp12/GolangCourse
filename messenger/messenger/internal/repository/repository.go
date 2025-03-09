package repository

import (
	"context"
	"messenger/internal/mto"
)

type Repository struct{
}


func (r Repository) CreateChat(context.Context, mto.Chat) (string, error){
	return "",nil
}

func (r Repository) CreateMessage(context.Context, mto.Message) (error){
	return nil
}
	
func (r Repository) GetMessagesFromChat(context.Context, string) ([]mto.Message, error){
	return nil,nil
}

