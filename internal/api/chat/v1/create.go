package server

import (
	"context"

	"github.com/andredubov/chat-server/internal/service/converter"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create is used to create a chat
func (c *Implementation) Create(ctx context.Context, r *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if r.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "chat name cannot be empty")
	}

	chatID, err := c.chatsService.Create(ctx, converter.ToChatFromCreateRequest(r))
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot add a new chat")
	}

	return &chat_v1.CreateResponse{
		Id: chatID,
	}, nil
}
