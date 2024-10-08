package converter

import (
	"github.com/andredubov/chat-server/internal/service/model"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
)

// ToMessageFromSendMessageRequest converts grpc request to message service layer model
func ToMessageFromSendMessageRequest(r *chat_v1.SendMessageRequest) model.Message {
	return model.Message{
		FromUserID: r.GetFromUserId(),
		ToChatID:   r.GetToChatId(),
		Text:       r.GetMessage(),
	}
}
