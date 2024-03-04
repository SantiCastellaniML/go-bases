package main

//import "bases/clases/03-funciones/internal/calculator"

//	"base/clases/03-funciones/internal/calculator"

import (
	"bases/clases/07-error-handling/app-2/internal/calculator"
	"errors"
)

func main() {
	a, b := 10, 0

	//division
	fnDiv, err := calculator.Orchestrator(calculator.OperatorDivide)

	if err != nil {
		println(err)
		return
	}

	result, err := fnDiv(a, b)
	if err != nil {
		//fmt.Println(err)
		//errUnwrapped := errors.Unwrap(err) //desencapsula el error. Esto es problemático porque no sabemos cuantos errores va a dar.
		//para resolver esto se hace una función recursiva, errors.Is

		switch {
		case errors.Is(err, calculator.ErrDivisionIndeterminate):
			println("Math error: division indeterminate")
		case errors.Is(err, calculator.ErrDivisionByZero):
			println("Math error: division by zero")
		default:
			println("Math error:", err)
		}

		/*
			switch errUnwrapped {
			case calculator.ErrDivisionIndeterminate:
				println("Math error: division indeterminate")
			case calculator.ErrDivisionByZero:
				println("Math error: division by zero")
			default:
				println("Math error:", err)
			}
		*/

		return
	}

	println(result)
}
