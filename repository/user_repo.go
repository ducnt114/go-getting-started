package repository

import (
	"context"
	"go-getting-started/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
	List(ctx context.Context, name string) ([]*model.User, error)
	Create(ctx context.Context, u *model.User) error
	FindByName(ctx context.Context, username string) (*model.User, error)
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
		Preload("Profile").
		Preload("Books").
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepo) Create(ctx context.Context, u *model.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepo) FindByName(ctx context.Context, username string) (*model.User, error) {
	return nil, nil // TODO
}

func (r *userRepo) List(ctx context.Context, name string) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("name LIKE ?", "%"+name+"%").
		Preload("Profile").
		Preload("Books").
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, err
}
