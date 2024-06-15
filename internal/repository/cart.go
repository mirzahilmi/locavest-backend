package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mirzahilmi/locavest-backend/internal/model"
)

type CartRepositoryItf interface {
	Create(ctx context.Context, banner *model.CartItem) error
	Read(ctx context.Context, size, page uint) ([]model.CartItem, error)
	Delete(ctx context.Context, slug string) error
}

type bannerRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) CartRepositoryItf {
	return &bannerRepository{db}
}

func (r *bannerRepository) Create(ctx context.Context, banner *model.CartItem) error {
	if _, err := r.db.NamedExecContext(ctx, "INSERT INTO cart (product_id, quantity) VALUES (?, ?)", *banner); err != nil {
		return err
	}
	return nil
}

func (r *bannerRepository) Delete(ctx context.Context, slug string) error {
	panic("unimplemented")
}

func (r *bannerRepository) Read(ctx context.Context, size uint, page uint) ([]model.CartItem, error) {
	panic("unimplemented")
}
