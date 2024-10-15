package main

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	chatService "github.com/avstriyskiy/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Create method
func (s *chatServer) Create(ctx context.Context, req *chatService.CreateRequest) (*chatService.CreateResponse, error) {
	log.Printf("creating chat with users: %v", req.GetUsernames())

	query, args, err := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("usernames").
		Values(req.GetUsernames()).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}
	log.Printf("query: %s, args: %v", query, args)

	var chatID int
	if err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID); err != nil {
		log.Printf("failed to create chat: %v, query: %s", err, query)
		return nil, err
	}

	log.Printf("successfully created chat with users: %v", req.GetUsernames())

	return &chatService.CreateResponse{
		Id: int64(chatID),
	}, nil
}

// SendMessage method
func (s *chatServer) SendMessage(ctx context.Context, req *chatService.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("sending message to chat from %s, message: %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}

// Delete method
func (s *chatServer) Delete(ctx context.Context, req *chatService.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("deleting chat ID: %d", req.GetId())

	query, args, err := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return nil, err
	}

	log.Printf("successfully deleted chat ID: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
