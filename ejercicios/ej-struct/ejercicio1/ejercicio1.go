package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products []Product = []Product{
	{ID: 1, Name: "Camisa", Price: 100.0, Description: "Camisa de algodon", Category: "Ropa"},
	{ID: 2, Name: "Pantalon", Price: 200.0, Description: "Pantalon de mezclilla", Category: "Ropa"},
	{ID: 3, Name: "Plancha", Price: 800.0},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) (Product, string) {
	for _, p := range Products {
		if p.ID == id {
			return p, ""
		}
	}

	return Product{}, fmt.Sprint("The product with the id ", id, " does not exist.")
}

func main() {
	p := Product{4, "Laptop", 2000.0, "Laptop de 16gb de ram", "Electronica"}
	p.GetAll()
	p.Save()
	p.GetAll()

	p2, err := getById(1)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(p2)
	}

	p3, err := getById(0)

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(p3)
	}
}
