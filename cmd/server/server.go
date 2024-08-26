package main

import (
	"github.com/avstriyskiy/chat-server/internal/config"
	chatService "github.com/avstriyskiy/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

type chatServer struct {
	chatService.UnimplementedChatServiceServer

	pool *pgxpool.Pool
	cfg  config.AppConfig
}
