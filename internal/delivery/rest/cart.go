package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mirzahilmi/locavest-backend/internal/model"
)

func RegisterCartHandler(api *echo.Group) {
	carts := api.Group("/carts")
	carts.GET("", fetchCartItems)
	carts.POST("", addToCart)
}

func fetchCartItems(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func addToCart(c echo.Context) error {
	var item model.CartItem
	if err := c.Bind(&item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}
