package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
)

func (p *participantsRepository) Create(ctx context.Context, chatID, userID int64) (int64, error) {
	builderInsert := sq.Insert(participantsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDParticipantsTableColumn, userIDParticipantsTableColumn).
		Values(chatID, userID).
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
