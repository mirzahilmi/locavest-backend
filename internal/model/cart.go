package model

type CartItem struct {
	ID        uint64 `db:"ID" json:"id"`
	ProductID uint64 `db:"ProductID" json:"productID"`
	Quantity  uint16 `db:"Quantity" json:"quantity"`
}
