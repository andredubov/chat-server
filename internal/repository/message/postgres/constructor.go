package postgres

import (
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/andredubov/golibs/pkg/client/database"
)

const (
	messagesTable              = "messages"
	idMessagesTableColumn      = "id"
	chatIDMessagesTableColumn  = "chat_id"
	userIDMessagesTableColumn  = "user_id"
	messageMessagesTableColumn = "text"
)

type messagesRepository struct {
	dbClient database.Client
}

// NewMessagesRepository create an instance of the messagesRepository struct
func NewMessagesRepository(dbClient database.Client) repository.Messages {
	return &messagesRepository{
		dbClient,
	}
}
