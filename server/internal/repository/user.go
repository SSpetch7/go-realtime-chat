package repository

import (
	"context"
	m "realtime_chat_server/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *m.User) (*m.User, error)
	GetUsers() (*[]m.User, error)
	GetUserByUsername(username string) (*m.User, error)
	GetUserByEmail(ctx context.Context, email string) (*m.User, error)
}
