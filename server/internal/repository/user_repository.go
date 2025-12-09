package repository

import (
	m "realtime_chat_server/internal/model"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(user *m.User) (*m.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (r userRepositoryDB) GetUsers() (users *[]m.User, err error) {
	return nil, nil
}
func (r userRepositoryDB) GetUserByID(id int64) (user *m.User, err error) {
	return nil, nil
}
