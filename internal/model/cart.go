package model

type CartItem struct {
	ID        uint64 `db:"ID" json:"id,omitempty"`
	ProductID uint64 `db:"ProductID" json:"productID,omitempty"`
	Quantity  uint16 `db:"Quantity" json:"quantity,omitempty"`
	Name      string `db:"Name" json:"name,omitempty"`
	Image     string `db:"Image" json:"image,omitempty"`
	Format    string `db:"Format" json:"format,omitempty"`
	Price     string `db:"Price" json:"price,omitempty"`
}
