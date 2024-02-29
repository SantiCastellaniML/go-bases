package ejercicio3

import "testing"

func TestCalcularSalario_A(t *testing.T) {
	//arange
	var minutos = 4 * 60
	var salario float32 = SalaryA
	var categoria = "A"

	var salarioEsperado float32 = float32(minutos) / 60.0 * salario * 1.5

	//act
	salarioObtenido, _ := CalcularSalario(minutos, categoria)

	//assert
	if salarioEsperado != salarioObtenido {
		t.Errorf("El salario esperado es %.2f, el salario obtenido es %.2f", salarioEsperado, salarioObtenido)
	}
}

func TestCalcularSalario_B(t *testing.T) {
	//arange
	var minutos = 4 * 60
	var salario float32 = SalaryB
	var categoria = "B"

	var salarioEsperado float32 = float32(minutos) / 60.0 * salario * 1.2

	//act
	salarioObtenido, _ := CalcularSalario(minutos, categoria)

	//assert
	if salarioEsperado != salarioObtenido {
		t.Errorf("El salario esperado es %.2f, el salario obtenido es %.2f", salarioEsperado, salarioObtenido)
	}
}

func TestCalcularSalario_C(t *testing.T) {
	//arange
	var minutos = 3 * 60
	var salario float32 = SalaryC
	var categoria = "C"

	var salarioEsperado float32 = float32(minutos) / 60.0 * salario

	//act
	salarioObtenido, _ := CalcularSalario(minutos, categoria)

	//assert
	if salarioEsperado != salarioObtenido {
		t.Errorf("El salario esperado es %.2f, el salario obtenido es %.2f", salarioEsperado, salarioObtenido)
	}
}
