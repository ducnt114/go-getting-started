package repository

import (
	"context"
	"go-getting-started/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
	List(ctx context.Context, name string) ([]*model.User, error)
	Create(ctx context.Context, u *model.User) error
	FindByName(ctx context.Context, username string) (*model.User, error)

	CreateUserWithBook(ctx context.Context, u *model.User) error
	UpdateUserAgeDemo(ctx context.Context, userID uint) (*model.User, error)
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

//func (r *userRepo) CreateUserWithBook(ctx context.Context, u *model.User) error {
//	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		if err := tx.Create(u).Error; err != nil {
//			return err // Rollback on error
//		}
//		// Insert book
//		for _, book := range u.Books {
//			if err := tx.Create(&book).Error; err != nil {
//				return err // Rollback on error
//			}
//		}
//		// Commit transaction if all operations succeed
//		return nil
//	})
//}

func (r *userRepo) CreateUserWithBook(ctx context.Context, u *model.User) error {
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback() // Rollback on error
		return err
	}
	tx.SavePoint("sp1")
	// Insert book
	for _, book := range u.Books {
		if err := tx.Create(&book).Error; err != nil {
			tx.RollbackTo("sp1")
			break
		}
	}
	tx.Commit() // Commit the transaction if all operations succeed
	return nil
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

func (r *userRepo) UpdateUserAgeDemo(ctx context.Context, userID uint) (*model.User, error) {
	user := &model.User{}
	tx := r.db.WithContext(ctx).Begin()

	// Fetch the product and lock the row for update
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	time.Sleep(10 * time.Second)

	// Update the stock
	user.Age += 10
	err := tx.Save(user).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return user, err
}
