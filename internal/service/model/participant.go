package model

import "time"

// Participant model for service layer
type Participant struct {
	ID        int64
	ChatID    int64
	UserID    int64
	CreatedAt time.Time
}
