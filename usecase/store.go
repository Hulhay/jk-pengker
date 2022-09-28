package usecase

import (
	"context"
	"strings"

	"github.com/Hulhay/jk-pengker/repository"
	"github.com/Hulhay/jk-pengker/usecase/store"
)

type storeUC struct {
	repo repository.StoreRepository
}

type Store interface {
	GetStoreList(ctx context.Context) ([]*store.StoreListResponse, error)
}

func NewStoreUC(r repository.StoreRepository) Store {
	return &storeUC{
		repo: r,
	}
}

func (u *storeUC) GetStoreList(ctx context.Context) ([]*store.StoreListResponse, error) {

	var res []*store.StoreListResponse

	stores, err := u.repo.GetStoreList(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range stores {
		res = append(res, &store.StoreListResponse{
			ID:       val.ID,
			Name:     val.Name,
			Category: strings.Split(val.Category, ","),
		})
	}

	return res, nil
}
