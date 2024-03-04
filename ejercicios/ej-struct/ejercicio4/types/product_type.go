package types

// constants for the types of products and errors
const (
	SmallProd     = "small"
	MediumProd    = "medium"
	LargeProd     = "large"
	ErrWrongType  = "The type doesn't exist"
	ErrWrongPrice = "The price is not a valid value"
)

type Tipo string

// product struct containing the category and the cost of the product, plus aditional data in a form of another struct
type Product struct {
	Category Tipo
	Cost     float32
	Aditionals
}

// aditional struct for products, containing aditional data such as if it's to be shipped or not
type Aditionals struct {
	Shipping bool
}

// interface for Products that can be sold
type Sellable interface {
	Price() float32
}

// default method for Product
func (p Product) Price() float32 {
	return 0
}
