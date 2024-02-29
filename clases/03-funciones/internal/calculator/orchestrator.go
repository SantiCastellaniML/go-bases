package calculator

func Orchestrator(operator MathOperator) (mo MathOperation, err string) {
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
		err = "Error: invalid operator"
	}

	return
}
