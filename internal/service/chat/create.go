package chat

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/model"
)

func (c *chatsService) Create(ctx context.Context, chat model.Chat) (int64, error) {
	var chatID int64
	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) (errTx error) {
		chatID, errTx = c.chatsRepository.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}

		for _, userID := range chat.UserIDs {
			participant := model.Participant{
				ChatID: chatID,
				UserID: userID,
			}

			_, errTx = c.participantsRepository.Create(ctx, participant)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}
