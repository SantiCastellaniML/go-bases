package ejercicio1

//si acá pongo package ejercicio1_test no me deja hacer import del package ejercicio1 para probar.

import (
	"testing"
)

func TestImpuestoDeSalario_debajoDe50000(t *testing.T) {
	//Arrange
	var salario float32 = 20000

	//Act
	result := ImpuestoDeSalario(salario)

	//Assert
	if result != salario {
		t.Errorf("El resultado deberia ser %f pero fue %f", salario, result)
	}
}

func TestImpuestoDeSalario_encimaDe50000(t *testing.T) {
	//Arrange
	var salario float32 = 100000
	var salarioEsperado float32 = salario * 0.83
	//Act
	result := ImpuestoDeSalario(salario)

	//Assert
	if result != salarioEsperado {
		t.Errorf("El resultado deberia ser %f pero fue %f", salarioEsperado, result)
	}
}

func TestImpuestoDeSalario_encimaDe150000(t *testing.T) {
	//Arrange
	var salario float32 = 200000
	var salarioEsperado float32 = salario * 0.73
	//Act
	result := ImpuestoDeSalario(salario)

	//Assert
	if result != salarioEsperado {
		t.Errorf("El resultado deberia ser %f pero fue %f", salarioEsperado, result)
	}
}
