package repository

import (
	"context"
	"go-getting-started/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error
	return &user, err
}
