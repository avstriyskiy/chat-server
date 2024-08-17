package main

import chatService "github.com/avstriyskiy/chat-server/pkg/chat_v1"

type chatServer struct {
	chatService.UnimplementedChatServiceServer
}
