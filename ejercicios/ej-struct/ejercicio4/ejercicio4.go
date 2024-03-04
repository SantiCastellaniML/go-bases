package main

import (
	"bases/ejercicios/ej-struct/ejercicio4/types"
	"errors"
	"fmt"
)

func main() {
	error1 := errors.New("hola que tal")

	fmt.Println(error1)

	p, err := Factory(types.LargeProd, 1000)

	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Println(p.Price())

	p, err = Factory(types.MediumProd, 1000)

	if err != "" {
		fmt.Println(err)
		//return
	}

	fmt.Println(p.Price())

	p, err = Factory(types.SmallProd, -1000)

	if err != "" {
		fmt.Println(err)
		//return
	}

	fmt.Println(p.Price())

	p, err = Factory("nada", 1000)

	if err != "" {
		fmt.Println(err)
		//return
	}

	fmt.Println(p.Price())
}

// Factory function for creating products of different types and prices given a type and a price.
// Returns a Sellable interface and an error string if the price is not valid or the type doesn't exist.
func Factory(t types.Tipo, price float32) (s types.Sellable, err string) {
	s = types.Product{Category: t, Cost: price}
	if price <= 0 {
		err = types.ErrWrongPrice
		return
	}

	switch t {
	case types.SmallProd:
		s = types.Small{Product: types.Product{Category: t, Cost: price}}
		return
	case types.MediumProd:
		s = types.Medium{Product: types.Product{Category: t, Cost: price}}
		return
	case types.LargeProd:
		s = types.Large{Product: types.Product{Category: t, Cost: price, Aditionals: types.Aditionals{Shipping: true}}}
		return
	default:
		err = types.ErrWrongType
		return
	}
}
