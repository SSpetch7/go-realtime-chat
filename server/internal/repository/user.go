package repository

import (
	m "realtime_chat_server/internal/model"
)

type UserRepository interface {
	CreateUser(user *m.User) (*m.User, error)
	GetUsers() (*[]m.User, error)
	GetUserByID(id int64) (*m.User, error)
}
