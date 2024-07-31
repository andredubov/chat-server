package server

import (
	"context"

	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Implementation) Delete(ctx context.Context, r *chat_v1.DeleteRequest) (*empty.Empty, error) {
	_, err := c.chatsService.Delete(ctx, r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot delete a chat")
	}

	return &empty.Empty{}, nil
}
