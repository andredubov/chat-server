package converter

import (
	"github.com/andredubov/chat-server/internal/service/model"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
)

// ToChatFromCreateRequest converts grpc request to chat service layer model
func ToChatFromCreateRequest(r *chat_v1.CreateRequest) model.Chat {
	return model.Chat{
		Name:    r.GetName(),
		UserIDs: r.GetUserIds(),
	}
}
