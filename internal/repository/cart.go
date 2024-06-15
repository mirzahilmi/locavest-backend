package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mirzahilmi/locavest-backend/internal/model"
)

type CartRepositoryItf interface {
	Create(ctx context.Context, cart *model.CartItem) error
	Read(ctx context.Context, size, page uint) ([]model.CartItem, error)
	Delete(ctx context.Context, slug string) error
}

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) CartRepositoryItf {
	return &cartRepository{db}
}

func (r *cartRepository) Create(ctx context.Context, cart *model.CartItem) error {
	if _, err := r.db.NamedExecContext(ctx, `
		INSERT INTO CartItems (ProductID, Quantity)
		VALUE (:ProductID, :Quantity);
	`, cart); err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) Delete(ctx context.Context, slug string) error {
	panic("unimplemented")
}

func (r *cartRepository) Read(ctx context.Context, size uint, page uint) ([]model.CartItem, error) {
	panic("unimplemented")
}
