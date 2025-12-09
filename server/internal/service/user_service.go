package service

import (
	"context"
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (s userService) Register(ctx context.Context, newUser *m.User) (*m.User, error) {
	user, err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
