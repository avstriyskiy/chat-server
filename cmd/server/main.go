package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/avstriyskiy/chat-server/internal/config"
	chatService "github.com/avstriyskiy/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Debug("Successfully loaded app config")

	postgresConfig, err := config.NewPostgresConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Debug("Successfully loaded postgres config")

	pool, err := pgxpool.Connect(ctx, postgresConfig.DSN())
	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Debug("Successfully connected to DB")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.GRPCPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chatService.RegisterChatServiceServer(s, &chatServer{
		cfg:  appConfig,
		pool: pool,
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
