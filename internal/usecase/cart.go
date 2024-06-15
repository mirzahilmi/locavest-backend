package usecase

import (
	"context"

	"github.com/mirzahilmi/locavest-backend/internal/model"
	"github.com/mirzahilmi/locavest-backend/internal/repository"
)

type cartUsecaseItf interface {
	Create(ctx context.Context, banner *model.CartItem) error
	Fetch(ctx context.Context) ([]model.CartItem, error)
	Delete(ctx context.Context, slug string) error
}

type bannerUsecase struct {
	bannerRepo repository.CartRepositoryItf
}

func NewcartUsecase(
	bannerRepo repository.CartRepositoryItf,
) cartUsecaseItf {
	return &bannerUsecase{bannerRepo}
}

func (u *bannerUsecase) Create(ctx context.Context, banner *model.CartItem) error {
	panic("unimplemented")
}

func (u *bannerUsecase) Delete(ctx context.Context, slug string) error {
	panic("unimplemented")
}

func (u *bannerUsecase) Fetch(ctx context.Context) ([]model.CartItem, error) {
	panic("unimplemented")
}
