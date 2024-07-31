package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
)

func (p *participantsRepository) Delete(ctx context.Context, participantID int64) (int64, error) {
	deleteBuilder := sq.Delete(participantsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idParticipantsTableColumn: participantID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, nil
	}

	q := database.Query{
		Name:     "participantsRepository.Delete",
		QueryRaw: query,
	}

	result, err := p.dbClient.Database().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
