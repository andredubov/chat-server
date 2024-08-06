package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
	"github.com/andredubov/chat-server/internal/service/model"
)

// Create a new chat in the chat repository
func (c *chatsRepository) Create(ctx context.Context, chat model.Chat) (int64, error) {
	insertBuilder := sq.Insert(chatsTable).PlaceholderFormat(sq.Dollar).
		Columns(nameChatsTableColumn).
		Values(chat.Name).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q := database.Query{
		Name:     "chatsRepository.Create",
		QueryRaw: query,
	}

	var chatID int64

	err = c.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
