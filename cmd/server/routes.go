package main

import (
	"context"
	"log"

	chatService "github.com/avstriyskiy/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Create method
func (s *chatServer) Create(_ context.Context, req *chatService.CreateRequest) (*chatService.CreateResponse, error) {
	log.Printf("Created chat with users: %v", req.GetUsernames())

	return &chatService.CreateResponse{
		Id: 888,
	}, nil
}

// SendMessage method
func (s *chatServer) SendMessage(_ context.Context, req *chatService.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("sending message to chat from %s, message: %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}

// Delete method
func (s *chatServer) Delete(_ context.Context, req *chatService.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting user ID: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
