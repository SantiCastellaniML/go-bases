package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	//arange
	a, b := 10, 5
	//act
	result, err := Divide(a, b)
	//assert
	if err != "" {
		t.Errorf("Expected no error but received error: %s", err)
		return
	}
	if result != 2 {
		t.Errorf("Expected result 2 but received %d", result)
		return
	}

	//assert.EqualValues(t, result, 2)
	//go test ./clases/03-funciones/internal/calculator
	require.Equal(t, result, 2)
}
