package main

import (
	"log"
	"net"

	"github.com/andredubov/chat-server/internal/config"
	"github.com/andredubov/chat-server/internal/config/env"
	server "github.com/andredubov/chat-server/internal/transport/grpc/chat/v1"
	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	listen, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatServer(s, server.NewChatServer())

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
