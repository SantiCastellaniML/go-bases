package main

import (
	"fmt"
)

const (
	salaryErrMin = "Error: the salary entered does not reach the taxable minimum"
)

type SalaryError struct {
	err string
}

func (e SalaryError) Error() string {
	return e.err
}

func main() {
	//Ejercicio 1
	var salary int = 40000

	_, err := ImpuestoDeSalario(float32(salary))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Must pay tax")
}

func ImpuestoDeSalario(salario float32) (float32, error) {
	if salario > 150000 {
		return salario * 0.73, nil
	}

	err := SalaryError{salaryErrMin}
	return 0, err
}
