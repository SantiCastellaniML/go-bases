package types

// Struct inheriting from Product with it's own defined Price method
type Small struct {
	Product
}

func (s Small) Price() float32 {
	return s.Cost
}
