package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	chatsTable               = "chats"
	idChatsTableColumn       = "id"
	nameChatsTableColumn     = "name"
	usersidsChatsTableColumn = "users_ids"
)

type chatsRepository struct {
	pool *pgxpool.Pool
}

// NewChatsRepository create an instance of the usersRepository struct
func NewChatsRepository(pool *pgxpool.Pool) repository.Chats {
	return &chatsRepository{
		pool,
	}
}

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

	var chatID int64

	err = c.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// Delete a chat from the chat repository
func (c *chatsRepository) Delete(ctx context.Context, chatID int64) (int64, error) {
	deleteBuilder := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idChatsTableColumn: chatID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, nil
	}

	result, err := c.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
