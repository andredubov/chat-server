package server

import (
	"github.com/andredubov/chat-server/internal/service"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
)

// Implementation ...
type Implementation struct {
	chat_v1.UnimplementedChatServer
	chatsService service.Chats
}

// NewImplementation creates new instance of Implementation struct
func NewImplementation(service service.Chats) *Implementation {
	return &Implementation{
		chatsService: service,
	}
}
