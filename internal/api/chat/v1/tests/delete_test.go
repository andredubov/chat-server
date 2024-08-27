package tests

import (
	"context"
	"testing"

	server "github.com/andredubov/chat-server/internal/api/chat/v1"
	"github.com/andredubov/chat-server/internal/service"
	serviceMocks "github.com/andredubov/chat-server/internal/service/mocks"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.Chats
	type args struct {
		ctx context.Context
		req *chat_v1.DeleteRequest
	}

	var (
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		id           = gofakeit.Int64()
		serviceError = status.Error(codes.Internal, "cannot delete a chat")

		request = &chat_v1.DeleteRequest{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *empty.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success service case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: &empty.Empty{},
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(id, nil)
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
				mock.DeleteMock.Expect(ctx, id).Return(0, serviceError)
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

			newID, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
