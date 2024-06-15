package usecase

import (
	"context"

	"github.com/mirzahilmi/locavest-backend/internal/model"
	"github.com/mirzahilmi/locavest-backend/internal/repository"
)

type CartUsecaseItf interface {
	Create(ctx context.Context, item *model.CartItem) error
	Fetch(ctx context.Context) ([]model.CartItem, error)
	Delete(ctx context.Context, id uint64) error
}

type cartUsecase struct {
	cartRepo repository.CartRepositoryItf
}

func NewcartUsecase(
	cartRepo repository.CartRepositoryItf,
) CartUsecaseItf {
	return &cartUsecase{cartRepo}
}

func (u *cartUsecase) Create(ctx context.Context, item *model.CartItem) error {
	if err := u.cartRepo.Create(ctx, item); err != nil {
		return err
	}
	return nil
}

func (u *cartUsecase) Delete(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

func (u *cartUsecase) Fetch(ctx context.Context) ([]model.CartItem, error) {
	panic("unimplemented")
}
