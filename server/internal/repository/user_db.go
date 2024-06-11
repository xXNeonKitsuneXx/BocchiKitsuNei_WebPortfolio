package repository

import (
	"bocchikitsunei_webportfolio/internal/entities"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAllUser() ([]entities.User, error) {
	users := []entities.User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r userRepositoryDB) GetUserByParams(userid int) (*entities.User, error) {
	users := entities.User{}
	result := r.db.Where("user_id = ?", userid).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
