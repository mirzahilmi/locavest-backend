package model

type CartItem struct {
	ID       uint64 `json:"id"`
	Quantity uint16 `json:"quantity"`
}
