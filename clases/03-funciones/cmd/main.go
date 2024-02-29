package main

import (
	"internal/calculator"
)

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
