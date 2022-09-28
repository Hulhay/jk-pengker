package repository

import (
	"context"

	"github.com/Hulhay/jk-pengker/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(ctx context.Context, params *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		qry: db,
	}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	if err := r.qry.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) InsertUser(ctx context.Context, params *model.User) error {
	var user *model.User

	if err := r.qry.Model(&user).Create(map[string]interface{}{
		"name":       params.Name,
		"email":      params.Email,
		"password":   params.Password,
		"role":       params.Role,
		"created_at": params.CreatedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}
