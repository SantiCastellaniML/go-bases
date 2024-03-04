package calculator

import "errors"

func Orchestrator(operator MathOperator) (mo MathOperation, err error) {
	switch operator {
	case OperatorSum:
		mo = Add
	case OperatorSubtract:
		mo = Subtract
	case OperatorMultiply:
		mo = Multiply
	case OperatorDivide:
		mo = Divide
	default:
		err = errors.New("Error: invalid operator")
	}

	return
}
