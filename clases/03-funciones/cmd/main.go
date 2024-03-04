package main

import "bases/clases/03-funciones/internal/calculator"

//	"base/clases/03-funciones/internal/calculator"

func main() {
	a, b := 10, 5

	//division
	fnDiv, err := calculator.Orchestrator(calculator.OperatorDivide)

	if err != "" {
		println(err)
		return
	}

	result, err := fnDiv(a, b)
	if err != "" {
		println(err)
		return
	}

	println(result)
}
