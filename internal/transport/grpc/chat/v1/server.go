package server

import (
	"context"
	"log"

	chat_v1 "github.com/andredubov/chat-server/pkg/chat/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type chatServer struct {
	chat_v1.UnimplementedChatServer
}

func NewChatServer() chat_v1.ChatServer {

	return &chatServer{}
}

func (c *chatServer) Create(ctx context.Context, r *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	const op = "ChatServer.GreateRequest"

	log.Printf("%s: input = %+v", op, r)

	return &chat_v1.CreateResponse{
		Id: 1,
	}, nil
}

func (c *chatServer) SendMessage(ctx context.Context, r *chat_v1.SendMessageRequest) (*empty.Empty, error) {

	const op = "ChatServer.SendMessageRequest"

	log.Printf("%s: input = %+v", op, r)

	return nil, nil
}

func (c *chatServer) Delete(ctx context.Context, r *chat_v1.DeleteRequest) (*empty.Empty, error) {

	const op = "ChatServer.DeleteRequest"

	log.Printf("%s: input = %+v", op, r)

	return nil, nil
}
