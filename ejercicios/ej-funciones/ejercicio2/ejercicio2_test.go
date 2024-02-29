package ejercicio2

import "testing"

func TestCalcularPromedio(t *testing.T) {
	//arange
	var notas = []float32{7, 10, 10, 5}
	var promedio float32 = 8

	//act
	result, err := CalcularPromedio(notas...)

	//assert
	//tuve que ejecutar con go test *.go, porque si no no encontraba la función calcularPromedio.
	if err != "" {
		t.Error("Expected no error but received error")
	}

	if result != promedio {
		t.Errorf("El promedio esperado es %.2f, el promedio obtenido es %.2f", promedio, result)
	}
}

func TestCalcularPromedio_notaNegativa(t *testing.T) {
	//arange
	var notas = []float32{7, 10, 10, -5, 5}

	//act
	_, err := CalcularPromedio(notas...)

	//assert
	//tuve que ejecutar con go test *.go, porque si no no encontraba la función calcularPromedio.
	if err == "" {
		t.Error("Expected error but received no error")
	}
}
