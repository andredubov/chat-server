package model

import "time"

type Message struct {
	ID         int64
	FromUserID int64
	ToChatID   int64
	Text       string
	CreatedAt  time.Time
}
