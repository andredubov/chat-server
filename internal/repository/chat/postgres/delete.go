package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/golibs/pkg/client/database"
)

// Delete a chat from the chat repository
func (c *chatsRepository) Delete(ctx context.Context, chatID int64) (int64, error) {
	deleteBuilder := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idChatsTableColumn: chatID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, nil
	}

	q := database.Query{
		Name:     "chatsRepository.Delete",
		QueryRaw: query,
	}

	result, err := c.dbClient.Database().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
