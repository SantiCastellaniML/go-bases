package main

import (
	"errors"
	"fmt"
)

const (
	salaryLess10000 = "Error: salary is less than 10000"
)

type SalaryError struct {
	err string
}

func (e SalaryError) Error() string {
	return salaryLess10000
}

var errSalaryLess10000 = SalaryError{}

//var errSalaryLess10000 = errors.New(salaryLess10000) //así sería para ejercicio 3. Sin el struct y sin la línea 20

func main() {
	//Ejercicio 2
	salario := 9999

	imp, err := ImpuestoDeSalario(float32(salario))

	if errors.Is(err, errSalaryLess10000) {
		fmt.Println(err)
		return
	}

	fmt.Println(imp)
}

func ImpuestoDeSalario(salario float32) (float32, error) {
	if salario > 150000 {
		return salario * 0.73, nil
	}

	if salario < 10000 {
		return 0, errSalaryLess10000 //errors.New(salaryLess10000) en ejercicio 3 sería así, para usar errors.New
	}

	return salario, nil
}
