package chat

import (
	"github.com/andredubov/chat-server/internal/repository"
	"github.com/andredubov/chat-server/internal/service"
	"github.com/andredubov/golibs/pkg/client/database"
)

type chatsService struct {
	chatsRepository        repository.Chats
	participantsRepository repository.Participants
	messagesRepository     repository.Messages
	txManager              database.TxManager
}

// NewService creates a new chats service that satisfies the service.Chats interface
func NewService(
	chatsRepository repository.Chats,
	participantsRepository repository.Participants,
	messagesRepository repository.Messages,
	txManager database.TxManager,
) service.Chats {
	return &chatsService{
		chatsRepository,
		participantsRepository,
		messagesRepository,
		txManager,
	}
}
