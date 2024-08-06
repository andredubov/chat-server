package server

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/converter"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SendMessage creates a new message in a chat
func (c *Implementation) SendMessage(ctx context.Context, r *chat_v1.SendMessageRequest) (*chat_v1.SendMessageResponse, error) {
	messageID, err := c.chatsService.SendMessage(ctx, converter.ToMessageFromSendMessageRequest(r))
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot create a message in chat")
	}

	return &chat_v1.SendMessageResponse{
		Id:     messageID,
		ChatId: r.GetToChatId(),
	}, nil
}
