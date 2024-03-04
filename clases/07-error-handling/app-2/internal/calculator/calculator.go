package calculator

import (
	"errors"
	"fmt"
)

type MathOperation func(int, int) (int, error)

type MathOperator string

var (
	ErrDivisionIndeterminate = errors.New("division indeterminate")
	ErrDivisionByZero        = errors.New("division by zero")
)

const (
	OperatorSum      MathOperator = "+"
	OperatorSubtract MathOperator = "-"
	OperatorMultiply MathOperator = "*"
	OperatorDivide   MathOperator = "/"
)

func Add(a, b int) (result int, err error) {
	result = a + b
	return
}

func Subtract(a, b int) (result int, err error) {
	result = a - b
	return
}

func Multiply(a, b int) (result int, err error) {
	result = a * b
	return
}

func Divide(a, b int) (result int, err error) {

	if b == 0 {
		if a == 0 {
			err = New(1000, a, b)
			err = fmt.Errorf("%w. %s", err, "Indeterminate division")
			return
		}

		err = New(1001, a, b)
		err = fmt.Errorf("%w. %s", err, "Division by zero")
		return
	}

	result = a / b
	return
}

// instancia de MathOperation:

/*
var Operation MathOperation = func(a, b int) int {
	return a + b
}
*/

func Calculate(a, b int, operation MathOperation) (int, error) {
	return operation(a, b)
}
