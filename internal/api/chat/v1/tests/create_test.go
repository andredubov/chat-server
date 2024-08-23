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

func TestCreateChat(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.Chats
	type args struct {
		ctx context.Context
		req *chat_v1.CreateRequest
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		id      = gofakeit.Int64()
		name    = gofakeit.Name()
		userIDs = []int64{4, 6, 8}

		serviceError = status.Error(codes.Internal, "cannot add a new chat")

		request = &chat_v1.CreateRequest{
			Name:    name,
			UserIds: userIDs,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *chat_v1.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "service success case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: &chat_v1.CreateResponse{
				Id: id,
			},
			err: nil,
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToChatFromCreateRequest(request)).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: nil,
			err:  serviceError,
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
				mock.CreateMock.Expect(ctx, converter.ToChatFromCreateRequest(request)).Return(0, serviceError)
				return mock
			},
		},

		{
			name: "empty chat name",
			args: args{
				ctx: ctx,
				req: &chat_v1.CreateRequest{
					Name:    "",
					UserIds: userIDs,
				},
			},
			want: nil,
			err:  status.Error(codes.InvalidArgument, "chat name cannot be empty"),
			chatServiceMock: func(mc *minimock.Controller) service.Chats {
				mock := serviceMocks.NewChatsMock(mc)
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

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
