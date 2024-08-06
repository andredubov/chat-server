package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
	"github.com/andredubov/chat-server/internal/service/model"
)

// Create is used to creates a new chat participants in the appropriate repository
func (p *participantsRepository) Create(ctx context.Context, participant model.Participant) (int64, error) {
	builderInsert := sq.Insert(participantsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDParticipantsTableColumn, userIDParticipantsTableColumn).
		Values(participant.ChatID, participant.UserID).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q, participantID := database.Query{
		Name:     "participantsRepository.Create",
		QueryRaw: query,
	}, int64(0)

	err = p.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&participantID)
	if err != nil {
		return 0, err
	}

	return participantID, nil
}
