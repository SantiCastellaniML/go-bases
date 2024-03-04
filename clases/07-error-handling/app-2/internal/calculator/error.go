package calculator

import "fmt"

type MathError struct {
	Code int
	Args []any
}

func New(code int, args ...any) error {
	return &MathError{
		Code: code,
		Args: args,
	}
}

func (e *MathError) Error() string {
	return fmt.Sprintf("Math error: code - %d, args - %v", e.Code, e.Args)
}
