package repository

import (
	"context"

	"github.com/Hulhay/jk-pengker/model"
	"gorm.io/gorm"
)

type StoreRepository interface {
	GetStoreList(ctx context.Context) ([]*model.Store, error)
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &repository{
		qry: db,
	}
}

func (r *repository) GetStoreList(ctx context.Context) ([]*model.Store, error) {
	var stores []*model.Store

	if err := r.qry.Model(&stores).Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}
