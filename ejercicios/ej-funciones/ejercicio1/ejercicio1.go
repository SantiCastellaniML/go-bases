package ejercicio1

import "fmt"

func main() {
	//Ejercicio 1
	fmt.Println(ImpuestoDeSalario(200000))
}

func ImpuestoDeSalario(salario float32) float32 {
	if salario > 150000 {
		return salario * 0.73
	}

	if salario > 50000 {
		return salario * 0.83
	}

	return salario
}
