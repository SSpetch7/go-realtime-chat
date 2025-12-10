package service

import (
	"context"
	"fmt"
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (s userService) Register(ctx context.Context, req *m.RegisterReq) (*m.User, error) {
	hashedPassword := ""

	// hashpassword

	newUser := &m.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	fmt.Println("check setup newUser : ", newUser)

	user, err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
