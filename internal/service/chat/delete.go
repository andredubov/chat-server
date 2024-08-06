package chat

import "context"

// Delete is used to delete a chat by its ID
func (c *chatsService) Delete(ctx context.Context, chatID int64) (int64, error) {
	quantity, err := c.chatsRepository.Delete(ctx, chatID)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
