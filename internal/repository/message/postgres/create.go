package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/service/model"
	"github.com/andredubov/golibs/pkg/client/database"
)

// Create a new message in the message repository
func (m *messagesRepository) Create(ctx context.Context, message model.Message) (int64, error) {
	builderInsert := sq.Insert(messagesTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDMessagesTableColumn, userIDMessagesTableColumn, messageMessagesTableColumn).
		Values(message.ToChatID, message.FromUserID, message.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := database.Query{
		Name:     "messagesRepository.Create",
		QueryRaw: query,
	}

	var messageID int64

	err = m.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&messageID)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}
