package repository

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/model"
)

// Chats defines inteface for chats repository
type Chats interface {
	Create(ctx context.Context, chat model.Chat) (int64, error)
	Delete(ctx context.Context, chatID int64) (int64, error)
}

// Messages defines inteface for message repository
type Messages interface {
	Create(ctx context.Context, message model.Message) (int64, error)
	Delete(ctx context.Context, messageID int64) (int64, error)
}

// Participants defines inteface for participans repository
type Participants interface {
	Create(ctx context.Context, participant model.Participant) (int64, error)
	Delete(ctx context.Context, participantID int64) (int64, error)
}
