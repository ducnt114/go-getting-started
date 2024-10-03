package service

import (
	"context"
	"go-getting-started/dto"
)

type UserService struct {
}

func (u *UserService) GetUserById(ctx context.Context, userId int64) (*dto.User, error) {
	// business logic here
	user, err := db.FindUserById(ctx, userId)
	res := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return nil, nil
}
