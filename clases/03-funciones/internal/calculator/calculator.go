package calculator

type MathOperation func(int, int) (int, string)

type MathOperator string

const (
	OperatorSum      MathOperator = "+"
	OperatorSubtract MathOperator = "-"
	OperatorMultiply MathOperator = "*"
	OperatorDivide   MathOperator = "/"
)

func Add(a, b int) (result int, error string) {
	result = a + b
	return
}

func Subtract(a, b int) (result int, error string) {
	result = a - b
	return
}

func Multiply(a, b int) (result int, error string) {
	result = a * b
	return
}

func Divide(a, b int) (result int, error string) {

	if b == 0 {
		return 0, "cannot divide by zero"
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

func Calculate(a, b int, operation MathOperation) (int, string) {
	return operation(a, b)
}
