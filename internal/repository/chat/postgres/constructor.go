package postgres

import (
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/andredubov/golibs/pkg/client/database"
)

const (
	chatsTable           = "chats"
	idChatsTableColumn   = "id"
	nameChatsTableColumn = "name"
)

type chatsRepository struct {
	dbClient database.Client
}

// NewChatsRepository create an instance of the usersRepository struct
func NewChatsRepository(dbClient database.Client) repository.Chats {
	return &chatsRepository{
		dbClient,
	}
}
