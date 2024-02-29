package ejercicio5

import "testing"

func TestAnimal_dog(t *testing.T) {
	//arange
	animalDog := "dog"
	var amount int = 30000

	//act
	animalFunc, err := Animal(animalDog)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := animalFunc(3)
	//assert
	if result != amount {
		t.Errorf("El amount esperado es %d, el amount obtenido es %d", amount, result)
		return
	}
}

func TestAnimal_cat(t *testing.T) {
	//arange
	animalCat := "cat"
	var amount int = 15000

	//act
	animalFunc, err := Animal(animalCat)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := animalFunc(3)
	//assert
	if result != amount {
		t.Errorf("El amount esperado es %d, el amount obtenido es %d", amount, result)
		return
	}
}
func TestAnimal_hamster(t *testing.T) {
	//arange
	animalHamster := hamster
	var amount int = hamsterFood * 3

	//act
	animalFunc, err := Animal(animalHamster)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := animalFunc(3)
	//assert
	if result != amount {
		t.Errorf("El amount esperado es %d, el amount obtenido es %d", amount, result)
		return
	}
}
func TestAnimal_tarantula(t *testing.T) {
	//arange
	animalTarantula := tarantula
	var amount int = tarantulaFood * 3

	//act
	animalFunc, err := Animal(animalTarantula)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := animalFunc(3)
	//assert
	if result != amount {
		t.Errorf("El amount esperado es %d, el amount obtenido es %d", amount, result)
		return
	}
}
