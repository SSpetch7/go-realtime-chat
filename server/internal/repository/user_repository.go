package repository

import (
	"context"
	"fmt"
	"log"
	m "realtime_chat_server/internal/model"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(ctx context.Context, user *m.User) (*m.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	res, err := r.GetUserByUsername(user.Username)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r userRepositoryDB) GetUsers() (users *[]m.User, err error) {
	return nil, nil
}

func (r userRepositoryDB) GetUserByUsername(username string) (user *m.User, err error) {
	result := r.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		log.Fatalf("Error get user : %v", result.Error)
	}

	return user, nil
}

func (r userRepositoryDB) GetUserByEmail(ctx context.Context, email string) (user *m.User, err error) {
	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("repo error : ", err)
		return user, nil
	}

	return user, nil
}
