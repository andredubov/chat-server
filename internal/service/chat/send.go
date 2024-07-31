package chat

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/model"
)

func (c *chatsService) SendMessage(ctx context.Context, message model.Message) (int64, error) {
	messageID, err := c.messagesRepository.Create(ctx, message)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}
