package service

import (
	"context"
	m "realtime_chat_server/internal/model"
)

type UserService interface {
	Register(ctx context.Context, newUser *m.RegisterReq) (*m.RegisterRes, error)
	Login(ctx context.Context, req *m.LoginReq) (*m.LoginRes, error)
}
