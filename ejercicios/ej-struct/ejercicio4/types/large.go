package types

// Struct inheriting from Product with it's own defined Price method
type Large struct {
	Product
}

func (l Large) Price() float32 {
	cost := l.Cost * 1.06

	if l.Shipping {
		cost += 2500
	}

	return cost
}
