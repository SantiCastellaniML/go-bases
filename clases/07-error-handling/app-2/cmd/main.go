package main

//import "bases/clases/03-funciones/internal/calculator"

//	"base/clases/03-funciones/internal/calculator"

import (
	"bases/clases/07-error-handling/app-2/internal/calculator"
	"errors"
)

func main() {
	a, b := 10, 0

	//division
	fnDiv, err := calculator.Orchestrator(calculator.OperatorDivide)

	if err != nil {
		println(err)
		return
	}

	result, err := fnDiv(a, b)
	if err != nil {
		//me interesa saber si el tipo de error que me llegó es MathError
		//errMath, ok := err.(*calculator.MathError) //esto ya no funciona si el error está encapsulado. Para eso se usa errors.As

		var errMath *calculator.MathError
		//el 2° valor es un puntero a una variable del tipo puntero del error que me interesa.
		ok := errors.As(err, &errMath) //donde voy a guardar el error si logra hacer el type assertion. Se debe pasar la dirección del puntero que es de ese tipo.

		//si yo le pasara un puntero en vez de un puntero al puntero: se pasaría una copia del puntero. Nunca pasaría a ser asignado ese puntero.
		//necesito pasarle el puntero del puntero para que pueda hacer el type assertion de ese puntero. Si le pasara solo la referencia de la variable sin ser puntero,
		//entonces el type asseretion se hace sobre la copia de esa dirección, y una vez que finaliza la función no pasa a ser de ese tipo el puntero (porque el assertion
		//se hace sobre la copia, no sobre el puntero original). Se busca modificar directamente el estado del puntero original, por eso se trabaja con una referencia al puntero original.

		//TODO ESTO DE ARRIBA SE HACE ASÍ PORQUE EN error.go EN New() Y EN EL MÉTODO Error() SE ESTÁ TRABAJANDO CON UN PUNTERO, ENTONCES NECESITO PASAR UN PUNTERO AL PUNTERO.
		//SI SE TRABAJARA DIRECTAMENTE CON LA VARIABLE, NO DEBERÍA HACER FALTA PASAR UN PUNTERO AL PUNTERO. PODRÍA PASARSE SOLAMENTE LA REFERENCIA DE LA VARIABLE.

		//2° preg: se queda con todo lo que seguiría de la cadena de errores? Si, el error no se ve afectado.

		if ok {
			switch errMath.Code {
			case 1000:
				println("Indeterminated division")
			case 1001:
				println("Division by zero")
			default:
				println("A math error occured:", err.Error())
			}

			return
		}

		println("Unknown error type", err.Error())
		return
	}

	println(result)
}
