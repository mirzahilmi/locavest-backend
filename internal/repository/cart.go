package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mirzahilmi/locavest-backend/internal/model"
)

type CartRepositoryItf interface {
	Create(ctx context.Context, cart *model.CartItem) error
	Read(ctx context.Context) ([]model.CartItem, error)
	Delete(ctx context.Context, id uint64) error
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

func (r *cartRepository) Read(ctx context.Context) ([]model.CartItem, error) {
	items := make([]model.CartItem, 0)
	if err := r.db.SelectContext(ctx, &items, `
		SELECT CI.ID, 
			P.Name,
			P.Image,
			P.Format,
			P.Price
		FROM CartItems CI
		INNER JOIN Products P ON CI.ProductID = P.ID
	`); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *cartRepository) Delete(ctx context.Context, id uint64) error {
	if _, err := r.db.ExecContext(ctx, `
		DELETE FROM CartItems WHERE ID = ?;
	`, id); err != nil {
		return err
	}
	return nil
}
