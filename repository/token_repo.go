package repository

import (
	"context"
	"github.com/samber/do"
	"go-getting-started/model"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Save(ctx context.Context, m *model.Token) error
	Delete(ctx context.Context, m *model.Token) error
	FindByRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error)
}

type tokenRepoImpl struct {
	db *gorm.DB
}

func newTokenRepository(di *do.Injector) (TokenRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &tokenRepoImpl{
		db: db,
	}, nil
}

func (r *tokenRepoImpl) Save(ctx context.Context, m *model.Token) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *tokenRepoImpl) Delete(ctx context.Context, m *model.Token) error {
	return r.db.WithContext(ctx).Delete(m).Error
}

func (r *tokenRepoImpl) FindByRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error) {
	var res model.Token
	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("refresh_token = ?", refreshToken).First(&res).Error
	return &res, err
}
