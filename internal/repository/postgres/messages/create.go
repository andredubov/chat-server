package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/client/database"
)

// Create a new message in the message repository
func (m *messagesRepository) Create(ctx context.Context, chatID, userID int64, message string) (int64, error) {
	builderInsert := sq.Insert(messagesTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDMessagesTableColumn, userIDMessagesTableColumn, messageMessagesTableColumn).
		Values(chatID, userID, message).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q, messageID := database.Query{
		Name:     "messagesRepository.Create",
		QueryRaw: query,
	}, int64(0)

	err = m.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&messageID)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}
