package model

type CartItem struct {
	ID        uint64 `db:"ID" json:"id"`
	ProductID uint64 `db:"ProductID" json:"productID,omitempty"`
	Quantity  uint16 `db:"Quantity" json:"quantity"`
	Name      string `db:"Name" json:"name"`
	Image     string `db:"Image" json:"image"`
	Format    string `db:"Format" json:"format"`
	Price     string `db:"Price" json:"price"`
}
