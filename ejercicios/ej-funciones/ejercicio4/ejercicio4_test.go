package ejercicio4

import "testing"

func TestOperation_min(t *testing.T) {
	//arange
	nums := []float32{1, 2, 3, 4, 5}
	var min float32 = 1

	//act
	fn, err := Operation(minimum)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := fn(nums)
	//assert
	if result != min {
		t.Errorf("El mínimo esperado es %.2f, el mínimo obtenido es %.2f", min, result)
		return
	}
}

func TestOperation_max(t *testing.T) {
	//arange
	nums := []float32{1, 2, 3, 4, 5}
	var max float32 = 5

	//act
	fn, err := Operation(maximum)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := fn(nums)
	//assert
	if result != max {
		t.Errorf("El máximo esperado es %.2f, el máximo obtenido es %.2f", max, result)
		return
	}
}

func TestOperation_avg(t *testing.T) {
	//arange
	nums := []float32{1, 2, 3, 4, 5}
	var avg float32 = 3

	//act
	fn, err := Operation(average)

	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}

	result := fn(nums)
	//assert
	if result != avg {
		t.Errorf("El promedio esperado es %.2f, el promedio obtenido es %.2f", avg, result)
		return
	}
}

func TestOperation_invalid(t *testing.T) {
	//arange
	var invalidOperation = "Hello"
	//act
	_, err := Operation(invalidOperation)

	//assert
	if err == "" {
		t.Errorf("Expected error but received no error")
		return
	}
}
