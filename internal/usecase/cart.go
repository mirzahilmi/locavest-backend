package usecase

import (
	"context"

	"github.com/mirzahilmi/locavest-backend/internal/model"
	"github.com/mirzahilmi/locavest-backend/internal/repository"
)

type CartUsecaseItf interface {
	Create(ctx context.Context, item *model.CartItem) error
	Checkout(ctx context.Context) error
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
	return u.cartRepo.Create(ctx, item)
}

func (u *cartUsecase) Checkout(ctx context.Context) error {
	return u.cartRepo.DeleteAll(ctx)
}

func (u *cartUsecase) Fetch(ctx context.Context) ([]model.CartItem, error) {
	return u.cartRepo.Read(ctx)
}

func (u *cartUsecase) Delete(ctx context.Context, id uint64) error {
	return u.cartRepo.Delete(ctx, id)
}
