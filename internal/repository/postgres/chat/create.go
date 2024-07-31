package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
)

// Create a new chat in the chat repository
func (c *chatsRepository) Create(ctx context.Context, name string) (int64, error) {
	insertBuilder := sq.Insert(chatsTable).PlaceholderFormat(sq.Dollar).
		Columns(nameChatsTableColumn).
		Values(name).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q, chatID := database.Query{
		Name:     "chatsRepository.Create",
		QueryRaw: query,
	}, int64(0)

	err = c.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
