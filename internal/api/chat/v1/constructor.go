package server

import (
	"github.com/andredubov/chat-server/internal/service"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
)

type Implementation struct {
	chat_v1.UnimplementedChatServer
	chatsService service.Chats
}

func NewImplementation(service service.Chats) *Implementation {
	return &Implementation{
		chatsService: service,
	}
}
