package repository

import (
	"context"
)

// Chats provide inteface for chats repository
type Chats interface {
	Create(ctx context.Context, name string, userIDs []int64) (int64, error)
	CreateMessage(ctx context.Context, userID, chatID int64, message string) (int64, int64, error)
	Delete(ctx context.Context, chatID int64) (int64, error)
}
