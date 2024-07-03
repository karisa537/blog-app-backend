package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/karisa537/blog-app/user/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByUsername implements UserRepository.
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}