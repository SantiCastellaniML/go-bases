package main

import "fmt"

func main() {
	//ejercicio 1:
	const nombre, direccion = "Santiago", "Hurlingham, Buenos Aires, Argentina"
	printDato("Nombre", nombre)
	printDato("Dirección", direccion)

	//ejercicio 2:
	var temperatura float32
	var humedad int //del 0 al 100%, si se quiere más decimales puede indicarse con un float en lugar de int
	var presion int //hPa enteros

	//ejercicio 3:
	//	var 1firstName string empieza con un número.
	var firstName string
	var lastName string
	//	var int age sería indicado al revés
	var age int
	//	1lastName := 6 empieza con un número y lastName debería ser un string.
	var lastName string
	var driver_license = true
	//	var person height int una variable debe ser una sola palabra, no puede contener espacios. El tipo de dato podría ser int si se representa en centímetros o float si es en metros.
	var personHeight int //puede ser int si se representa en centímetros o float si es en metros.
	childsNumber := 2

	//ejercicio 4:
	var lastName string = "Smith"
	var age int = "35"           //va sin comillas.
	boolean := "false"           //va sin comillas y la variable no puede llamarse como el primitivo.
	var salary string = 45857.90 //si es string va con comillas, si fuera float en vez de string estaría bien.
	var firstName string = "Mary"
}

func printDato(clave string, valor string) {
	if clave == "" {
		fmt.Println("No se ingresó nombre del campo.")
	} else if valor == "" {
		fmt.Println("No se ingresó el valor del dato.")
	} else {
		fmt.Println(clave + ": " + valor)
	}
}
