package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/golibs/pkg/client/database"
)

// Delete a message from the message repository
func (m *messagesRepository) Delete(ctx context.Context, messageID int64) (int64, error) {
	deleteBuilder := sq.Delete(messagesTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idMessagesTableColumn: messageID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, nil
	}

	q := database.Query{
		Name:     "messagesRepository.Delete",
		QueryRaw: query,
	}

	result, err := m.dbClient.Database().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
