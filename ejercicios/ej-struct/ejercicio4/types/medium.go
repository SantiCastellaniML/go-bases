package types

// Struct inheriting from Product with it's own defined Price method
type Medium struct {
	Product
}

func (m Medium) Price() float32 {
	return m.Cost * 1.03
}
