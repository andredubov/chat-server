package tests

import (
	"context"
	"testing"

	server "github.com/andredubov/chat-server/internal/api/chat/v1"
	"github.com/andredubov/chat-server/internal/service"
	"github.com/andredubov/chat-server/internal/service/converter"
	serviceMocks "github.com/andredubov/chat-server/internal/service/mocks"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.Chats
	type args struct {
		ctx context.Context
		req *chat_v1.SendMessageRequest
	}

	var (
		ctx       = context.Background()
		mc        = minimock.NewController(t)
		userID    = gofakeit.Int64()
		chatID    = gofakeit.Int64()
		messageID = gofakeit.Int64()
		message   = gofakeit.Name()

		serviceError = status.Error(codes.Internal, "cannot create a message in chat")

		request = &chat_v1.SendMessageRequest{
			FromUserId: userID,
			ToChatId:   chatID,
			Message:    message,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *chat_v1.SendMessageResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success service case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: &chat_v1.SendMessageResponse{
				Id:     messageID,
				ChatId: chatID,
			},
			err: nil,
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
				mock.SendMessageMock.Expect(ctx, converter.ToMessageFromSendMessageRequest(request)).Return(messageID, nil)
				return mock
			},
		},
		{
			name: "error service case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: nil,
			err:  serviceError,
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
				mock.SendMessageMock.Expect(ctx, converter.ToMessageFromSendMessageRequest(request)).Return(0, serviceError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := server.NewImplementation(chatServiceMock)

			newID, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
