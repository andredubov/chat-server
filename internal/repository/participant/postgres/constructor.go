package postgres

import (
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/andredubov/golibs/pkg/client/database"
)

const (
	participantsTable             = "participants"
	idParticipantsTableColumn     = "id"
	chatIDParticipantsTableColumn = "chat_id"
	userIDParticipantsTableColumn = "user_id"
)

type participantsRepository struct {
	dbClient database.Client
}

// NewParticipantsRepository create an instance of the participantsRepository struct
func NewParticipantsRepository(dbClient database.Client) repository.Participants {
	return &participantsRepository{
		dbClient,
	}
}
