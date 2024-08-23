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

func TestCreateChat(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.Chats
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.Messages
	type participantRepositoryMockFunc func(mc *minimock.Controller) repository.Participants
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx   context.Context
		input model.Chat
	}

	var (
		ctx  = context.Background()
		mc   = minimock.NewController(t)
		id   = gofakeit.Int64()
		ids  = []int64{2, 3, 4}
		name = gofakeit.Name()

		repositoryError = errors.New("repo error")

		chat = model.Chat{
			ID:      id,
			Name:    name,
			UserIDs: ids,
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
				ctx:   ctx,
				input: chat,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.Chats {
				mock := repoMocks.NewChatsMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(id, nil)
				return mock
			},
			participantRepositoryMock: func(mc *minimock.Controller) repository.Participants {
				mock := repoMocks.NewParticipantsMock(mc)
				for _, userID := range chat.UserIDs {
					participant := model.Participant{
						ChatID: chat.ID,
						UserID: userID,
					}
					mock.CreateMock.When(ctx, participant).Then(0, nil)
				}
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.Messages {
				mock := repoMocks.NewMessagesMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f database.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "error case (chat repo error)",
			args: args{
				ctx:   ctx,
				input: chat,
			},
			want: 0,
			err:  repositoryError,
			chatRepositoryMock: func(mc *minimock.Controller) repository.Chats {
				mock := repoMocks.NewChatsMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(0, repositoryError)
				return mock
			},
			participantRepositoryMock: func(mc *minimock.Controller) repository.Participants {
				mock := repoMocks.NewParticipantsMock(mc)
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.Messages {
				mock := repoMocks.NewMessagesMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f database.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "error case (participant repo error)",
			args: args{
				ctx:   ctx,
				input: chat,
			},
			want: 0,
			err:  repositoryError,
			chatRepositoryMock: func(mc *minimock.Controller) repository.Chats {
				mock := repoMocks.NewChatsMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(id, nil)
				return mock
			},
			participantRepositoryMock: func(mc *minimock.Controller) repository.Participants {
				mock := repoMocks.NewParticipantsMock(mc)
				for _, userID := range chat.UserIDs {
					participant := model.Participant{
						ChatID: chat.ID,
						UserID: userID,
					}
					mock.CreateMock.Expect(ctx, participant).Return(0, repositoryError)
					break
				}
				return mock
			},
			messageRepositoryMock: func(mc *minimock.Controller) repository.Messages {
				mock := repoMocks.NewMessagesMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f database.Handler) error {
					return f(ctx)
				})
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

			newID, err := service.Create(tt.args.ctx, tt.args.input)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
