package service

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/model"
)

// Chats interface for working with a chats service
type Chats interface {
	Create(ctx context.Context, chat model.Chat) (int64, error)
	SendMessage(ctx context.Context, message model.Message) (int64, error)
	Delete(ctx context.Context, chatID int64) (int64, error)
}
