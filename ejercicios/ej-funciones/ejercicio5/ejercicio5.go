package ejercicio5

import "fmt"

const (
	dogFood       = 10000
	catFood       = 5000
	hamsterFood   = 250
	tarantulaFood = 150
	dog           = "dog"
	cat           = "cat"
	hamster       = "hamster"
	tarantula     = "tarantula"
)

func main() {
	animalDog, msg := Animal(dog)

	if msg != "" {
		println(msg)
		return
	}

	animalCat, msg := Animal(cat)

	if msg != "" {
		println(msg)
		return
	}

	var amount int

	amount += animalDog(10)

	amount += animalCat(10)

	fmt.Println("Total amount: ", float32(amount/1000), "Kg")
}

func Animal(animal string) (func(int) int, string) {
	var animalFunc func(int) int
	switch animal {
	case dog:
		animalFunc = dogAmount
	case cat:
		animalFunc = catAmount
	case hamster:
		animalFunc = hamsterAmount
	case tarantula:
		animalFunc = tarantulaAmount
	default:
		return nil, "Error: cannot find animal"
	}

	return animalFunc, ""
}

func dogAmount(amount int) int {
	return amount * dogFood
}

func catAmount(amount int) int {
	return amount * catFood
}

func hamsterAmount(amount int) int {
	return amount * hamsterFood
}

func tarantulaAmount(amount int) int {
	return amount * tarantulaFood
}
