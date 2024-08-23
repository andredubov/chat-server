package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/andredubov/chat-server/internal/repository"
	repoMocks "github.com/andredubov/chat-server/internal/repository/mocks"
	serviceMocks "github.com/andredubov/chat-server/internal/service/chat"
	"github.com/andredubov/chat-server/internal/service/model"
	"github.com/andredubov/golibs/pkg/client/database"
	dbMocks "github.com/andredubov/golibs/pkg/client/database/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
)

func TestSendChat(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.Chats
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.Messages
	type participantRepositoryMockFunc func(mc *minimock.Controller) repository.Participants
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx     context.Context
		message model.Message
	}

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		repositoryError = errors.New("repo error")
		id              = gofakeit.Int64()
		userID          = gofakeit.Int64()
		chatID          = gofakeit.Int64()
		text            = gofakeit.Name()

		message = model.Message{
			ID:         id,
			FromUserID: userID,
			ToChatID:   chatID,
			Text:       text,
		}
	)

	tests := []struct {
		name                      string
		args                      args
		want                      int64
		err                       error
		chatRepositoryMock        chatRepositoryMockFunc
		participantRepositoryMock participantRepositoryMockFunc
		messageRepositoryMock     messageRepositoryMockFunc
		txManagerMock             txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:     ctx,
				message: message,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.Chats {
				mock := repoMocks.NewChatsMock(mc)
				return mock
			},
			participantRepositoryMock: func(mc *minimock.Controller) repository.Participants {
				mock := repoMocks.NewParticipantsMock(mc)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.Messages {
				mock := repoMocks.NewMessagesMock(mc)
				mock.CreateMock.Expect(ctx, message).Return(id, nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:     ctx,
				message: message,
			},
			want: 0,
			err:  repositoryError,
			chatRepositoryMock: func(mc *minimock.Controller) repository.Chats {
				mock := repoMocks.NewChatsMock(mc)
				return mock
			},
			participantRepositoryMock: func(mc *minimock.Controller) repository.Participants {
				mock := repoMocks.NewParticipantsMock(mc)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.Messages {
				mock := repoMocks.NewMessagesMock(mc)
				mock.CreateMock.Expect(ctx, message).Return(0, repositoryError)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			chatRepositoryMock := tt.chatRepositoryMock(mc)
			participantRepositoryMock := tt.participantRepositoryMock(mc)
			messageRepositoryMock := tt.messageRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := serviceMocks.NewService(
				chatRepositoryMock,
				participantRepositoryMock,
				messageRepositoryMock,
				txManagerMock,
			)

			newID, err := service.SendMessage(tt.args.ctx, tt.args.message)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
