package repository

import (
	"context"
)

// Chats provide inteface for chats repository
type Chats interface {
	Create(ctx context.Context, name string, userIDs []int64) (int64, error)
	Delete(ctx context.Context, chatID int64) (int64, error)
}

// Messages provide inteface for message repository
type Messages interface {
	Create(ctx context.Context, chatID, userID int64, message string) (int64, error)
	Delete(ctx context.Context, messageID int64) (int64, error)
}
