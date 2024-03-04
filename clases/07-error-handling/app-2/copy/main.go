package copy

import "fmt"

func New(text string) error {
	return CustomError{text}
}

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

var (
	ErrRandom = New("Something random happened")
)

func main() {
	if ErrRandom == New("Something random happened") { //esto da verdadero porque los campos de la estructura son iguales. Para comparar errores se deben comparar sus punteros, no sus contenidos.
		fmt.Println("Errors are equal")
		return
	}

	fmt.Println("Error messages are different")
}
