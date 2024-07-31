package postgres

import (
	"github.com/andredubov/chat-server/internal/client/database"
	"github.com/andredubov/chat-server/internal/repository"
)

const (
	chatsTable           = "chats"
	idChatsTableColumn   = "id"
	nameChatsTableColumn = "name"
)

type chatsRepository struct {
	dbClient database.Client
}

// NewUsersRepository create an instance of the usersRepository struct
func NewUsersRepository(dbClient database.Client) repository.Chats {
	return &chatsRepository{
		dbClient,
	}
}
