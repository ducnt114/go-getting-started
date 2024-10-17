package repository

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"go-getting-started/model"
	"gorm.io/gorm"
	"runtime"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
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
	span := sentry.StartSpan(ctx, "userRepo.FindByID")
	span.Description = "GetUserById_repo"
	defer span.Finish()

	//time.Sleep(1 * time.Second)

	var user model.User
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error

	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	printStack()
	//}
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepo) Create(ctx context.Context, u *model.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func printStack() {
	// Create a slice to store the program counters
	pc := make([]uintptr, 10)
	// Skip 2 frames: runtime.Callers and printStack itself
	n := runtime.Callers(2, pc)
	// Retrieve and format the call stack frames
	frames := runtime.CallersFrames(pc[:n])
	// Iterate over the frames and print information
	for {
		frame, more := frames.Next()
		fmt.Printf("Function: %s\nFile: %s\nLine: %d\n\n",
			frame.Function, frame.File, frame.Line)
		// Check if there are more frames
		if !more {
			break
		}
	}
}

func (r *userRepo) FindByName(ctx context.Context, username string) (*model.User, error) {
	return nil, nil // TODO
}
