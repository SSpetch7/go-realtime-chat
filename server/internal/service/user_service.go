package service

import (
	"context"
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/repository"
	"realtime_chat_server/util"
	"strconv"
	"time"
)

type userService struct {
	userRepo repository.UserRepository
	timeout  time.Duration
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo, time.Duration(2) * time.Second}
}

func (s userService) Register(c context.Context, req *m.RegisterReq) (*m.RegisterRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)

	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	newUser := &m.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	res, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user := m.RegisterRes{
		ID:       strconv.Itoa(int(res.ID)),
		Username: res.Username,
		Email:    res.Email,
	}

	return &user, nil
}
