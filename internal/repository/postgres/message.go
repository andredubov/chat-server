package postgres

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	messagesTable = "messages"
)

type messagesRepository struct {
	pool *pgxpool.Pool
}

// NewMessagesRepository create an instance of the messagesRepository struct
func NewMessagesRepository(pool *pgxpool.Pool) repository.Messages {
	return &messagesRepository{
		pool,
	}
}

// Create a new message in the message repository
func (m *messagesRepository) Create(ctx context.Context, chatID, userID int64, message string) (int64, error) {
	const op = "messagesRepository.Create"

	builderInsert := sq.Insert(messagesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(chatID, userID, message).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, err
	}

	var messageID int64

	err = m.pool.QueryRow(ctx, query, args...).Scan(&messageID)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, err
	}

	return messageID, nil
}

// Delete a message from the message repository
func (m *messagesRepository) Delete(ctx context.Context, messageID int64) (int64, error) {
	const op = "messagesRepository.Delete"

	deleteBuilder := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": messageID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, nil
	}

	result, err := m.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, err
	}

	return result.RowsAffected(), nil
}
