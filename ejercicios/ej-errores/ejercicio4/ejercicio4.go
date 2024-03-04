package main

import (
	"fmt"
)

const (
	salaryLess10000 = "Error: salary is less than 10000"
)

func main() {
	//Ejercicio 4
	var salario float32 = 149999

	//errSalaryLowerThanMin := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", salario)

	imp, err := ImpuestoDeSalario(float32(salario))

	//if errors.Is(err, errSalaryLowerThanMin) { can't do it this way because the error is not predefined. It's a variable string because of it's parameter.
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return
	}

	fmt.Println(imp)
}

func ImpuestoDeSalario(salario float32) (float32, error) {
	var errSalaryLowerThanMin = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", salario)

	if salario > 150000 {
		return salario * 0.73, nil
	}

	return 0, errSalaryLowerThanMin
}
