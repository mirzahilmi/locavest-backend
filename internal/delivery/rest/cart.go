package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mirzahilmi/locavest-backend/internal/model"
	"github.com/mirzahilmi/locavest-backend/internal/usecase"
)

type cartHandler struct {
	usecase usecase.CartUsecaseItf
}

func RegisterCartHandler(api *echo.Group, usecase usecase.CartUsecaseItf) {
	carts := api.Group("/carts")
	handler := cartHandler{usecase}
	carts.GET("", handler.fetchCartItems)
	carts.POST("", handler.addToCart)
	carts.DELETE("/:id", handler.removeCartItem)
}

func (h *cartHandler) fetchCartItems(c echo.Context) error {
	items, err := h.usecase.Fetch(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (h *cartHandler) addToCart(c echo.Context) error {
	var item model.CartItem
	if err := c.Bind(&item); err != nil {
		return err
	}
	if err := h.usecase.Create(c.Request().Context(), &item); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *cartHandler) removeCartItem(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}
	if err := h.usecase.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
