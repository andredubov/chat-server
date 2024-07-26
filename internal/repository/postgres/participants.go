package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	participantsTable             = "participants"
	idParticipantsTableColumn     = "id"
	chatIDParticipantsTableColumn = "chat_id"
	userIDParticipantsTableColumn = "user_id"
)

type participantsRepository struct {
	pool *pgxpool.Pool
}

// NewParticipantsRepository create an instance of the participantsRepository struct
func NewParticipantsRepository(pool *pgxpool.Pool) repository.Participants {
	return &participantsRepository{
		pool,
	}
}

func (p *participantsRepository) Create(ctx context.Context, chatID, userID int64) (int64, error) {
	builderInsert := sq.Insert(participantsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDMessagesTableColumn, userIDMessagesTableColumn).
		Values(chatID, userID).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var participantID int64

	err = p.pool.QueryRow(ctx, query, args...).Scan(&participantID)
	if err != nil {
		return 0, err
	}

	return participantID, nil
}

func (p *participantsRepository) Delete(ctx context.Context, participantID int64) (int64, error) {
	deleteBuilder := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idParticipantsTableColumn: participantID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, nil
	}

	result, err := p.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
