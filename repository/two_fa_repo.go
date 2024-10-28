package repository

import (
	"context"
	"github.com/samber/do"
	"go-getting-started/model"
	"gorm.io/gorm"
)

type TwoFaRepository interface {
	Save(ctx context.Context, m *model.TwoFactor) error
	Delete(ctx context.Context, m *model.TwoFactor) error
	FindByUserID(ctx context.Context, userID uint) (*model.TwoFactor, error)
}

type twoFaRepoImpl struct {
	db *gorm.DB
}

func newTwoFaRepository(di *do.Injector) (TwoFaRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &twoFaRepoImpl{
		db: db,
	}, nil
}

func (r *twoFaRepoImpl) Save(ctx context.Context, m *model.TwoFactor) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *twoFaRepoImpl) Delete(ctx context.Context, m *model.TwoFactor) error {
	return r.db.WithContext(ctx).Delete(m).Error
}

func (r *twoFaRepoImpl) FindByUserID(ctx context.Context, userID uint) (*model.TwoFactor, error) {
	var res model.TwoFactor
	err := r.db.WithContext(ctx).Model(&model.TwoFactor{}).
		Where("user_id = ?", userID).First(&res).Error
	return &res, err
}
