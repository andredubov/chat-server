package app

import (
	"context"
	"log"

	server "github.com/andredubov/chat-server/internal/api/chat/v1"
	"github.com/andredubov/chat-server/internal/repository"
	chatsRepo "github.com/andredubov/chat-server/internal/repository/chat/postgres"
	messagesRepo "github.com/andredubov/chat-server/internal/repository/message/postgres"
	participantRepo "github.com/andredubov/chat-server/internal/repository/participant/postgres"
	"github.com/andredubov/chat-server/internal/service"
	"github.com/andredubov/chat-server/internal/service/chat"
	"github.com/andredubov/golibs/pkg/client/database"
	postgresClient "github.com/andredubov/golibs/pkg/client/database/postgres"
	"github.com/andredubov/golibs/pkg/client/database/transaction"
	"github.com/andredubov/golibs/pkg/closer"
	"github.com/andredubov/golibs/pkg/config"
	"github.com/andredubov/golibs/pkg/config/env"
)

type serviceProvider struct {
	postgresConfig         config.PostgresConfig
	grpcConfig             config.GRPCConfig
	databaseClient         database.Client
	databaseTxManager      database.TxManager
	chatsRepository        repository.Chats
	messagesRepository     repository.Messages
	participantsRepository repository.Participants
	chatsService           service.Chats
	serverImplementation   *server.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PostgresConfig loads postgres config from an appropriate enviroment variables
func (s *serviceProvider) PostgresConfig() config.PostgresConfig {
	if s.postgresConfig == nil {
		cfg, err := env.NewPostgresConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		s.postgresConfig = cfg
	}

	return s.postgresConfig
}

// GRPCConfig loads grpc server config from an appropriate enviroment variables
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// DatabaseClient creates a database client
func (s *serviceProvider) DatabaseClient(ctx context.Context) database.Client {
	if s.databaseClient == nil {
		dbClient, err := postgresClient.New(ctx, s.PostgresConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := dbClient.Database().Ping(ctx); err != nil {
			log.Fatalf("database ping error: %v", err)
		}

		closer.Add(func() error {
			dbClient.Database().Close()
			return nil
		})

		s.databaseClient = dbClient
	}

	return s.databaseClient
}

// TxManager creates an instance of a transaction manager
func (s *serviceProvider) TxManager(ctx context.Context) database.TxManager {
	if s.databaseTxManager == nil {
		db := s.DatabaseClient(ctx).Database()
		s.databaseTxManager = transaction.NewTransactionManager(db)
	}

	return s.databaseTxManager
}

// ChatsRepository creates an insanse of a chata repository
func (s *serviceProvider) ChatsRepository(ctx context.Context) repository.Chats {
	if s.chatsRepository == nil {
		dbClient := s.DatabaseClient(ctx)
		s.chatsRepository = chatsRepo.NewChatsRepository(dbClient)
	}

	return s.chatsRepository
}

// MessagesRepository creates an instance of messages repository
func (s *serviceProvider) MessagesRepository(ctx context.Context) repository.Messages {
	if s.messagesRepository == nil {
		dbClient := s.DatabaseClient(ctx)
		s.messagesRepository = messagesRepo.NewMessagesRepository(dbClient)
	}

	return s.messagesRepository
}

// ParticipantsRepository creates an instance of a participants repository
func (s *serviceProvider) ParticipantsRepository(ctx context.Context) repository.Participants {
	if s.participantsRepository == nil {
		dbClient := s.DatabaseClient(ctx)
		s.participantsRepository = participantRepo.NewParticipantsRepository(dbClient)
	}

	return s.participantsRepository
}

// ChatsService creates an instance of chats service
func (s *serviceProvider) ChatsService(ctx context.Context) service.Chats {
	if s.chatsService == nil {
		s.chatsService = chat.NewService(
			s.ChatsRepository(ctx),
			s.ParticipantsRepository(ctx),
			s.MessagesRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatsService
}

// ServerImplementation creates an instance of grpc server implementation
func (s *serviceProvider) ServerImplementation(ctx context.Context) *server.Implementation {
	if s.serverImplementation == nil {
		chatsService := s.ChatsService(ctx)
		s.serverImplementation = server.NewImplementation(chatsService)
	}

	return s.serverImplementation
}
