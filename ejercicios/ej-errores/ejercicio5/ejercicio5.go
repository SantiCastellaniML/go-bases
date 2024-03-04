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
	//Ejercicio 5

	var hours int = 140
	var hourValue float32 = 1500

	salary, err := monthSalary(hours, hourValue)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Salary: ", salary, "\nSalary minus taxes: ", ImpuestoDeSalario(salary))
}

func monthSalary(hours int, hourValue float32) (salary float32, err error) {
	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	return float32(hours) * hourValue, nil
}

func ImpuestoDeSalario(salario float32) float32 {
	if salario > 150000 {
		return salario * 0.9
	}

	return salario
}
