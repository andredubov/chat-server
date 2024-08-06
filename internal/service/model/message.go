package model

import "time"

// Message model for service layer
type Message struct {
	ID         int64
	FromUserID int64
	ToChatID   int64
	Text       string
	CreatedAt  time.Time
}
