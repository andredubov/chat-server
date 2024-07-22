package server

import (
	"context"

	"github.com/andredubov/chat-server/internal/repository"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type chatServer struct {
	chat_v1.UnimplementedChatServer
	chatsRepository    repository.Chats
	messagesRepository repository.Messages
}

// NewChatServer returns an instance of charServer struct
func NewChatServer(chatsRepo repository.Chats, messagesRepo repository.Messages) chat_v1.ChatServer {
	return &chatServer{
		chatsRepository:    chatsRepo,
		messagesRepository: messagesRepo,
	}
}

func (c *chatServer) Create(ctx context.Context, r *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if r.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "chat name cannot be empty")
	}

	chatID, err := c.chatsRepository.Create(ctx, r.GetName(), r.GetUserIds())
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot add a new chat")
	}

	return &chat_v1.CreateResponse{
		Id: chatID,
	}, nil
}

func (c *chatServer) Delete(ctx context.Context, r *chat_v1.DeleteRequest) (*empty.Empty, error) {
	_, err := c.chatsRepository.Delete(ctx, r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot delete a chat")
	}

	return &empty.Empty{}, nil
}

func (c *chatServer) SendMessage(ctx context.Context, r *chat_v1.SendMessageRequest) (*chat_v1.SendMessageResponse, error) {
	messageID, err := c.messagesRepository.Create(ctx, r.GetToChatId(), r.GetFromUserId(), r.GetMessage())
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot add chat message")
	}

	return &chat_v1.SendMessageResponse{
		Id:     messageID,
		ChatId: r.GetToChatId(),
	}, nil
}
