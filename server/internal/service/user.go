package service

import (
	"context"
	m "realtime_chat_server/internal/model"
)

type UserService interface {
	Register(ctx context.Context, newUser *m.User) (*m.User, error)
}
