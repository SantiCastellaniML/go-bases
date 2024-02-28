package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Ejercicio 1: Letras de una palabra
	procesarPalabra("Hola")

	//Ejercicio 2: Préstamo
	edad, empleado, antiguedad, sueldo := 21, false, 3, 99999
	determinarPrestamo(edad, empleado, antiguedad, sueldo)

	//Ejercicio 3: A qué mes corresponde
	mes := 3
	fmt.Println(procesarMes(mes))
	fmt.Println(procesarMesConSwitch(mes))

	//Ejercicio 4: Qué edad tiene...
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	nombre := "Benjamin"
	edadNombre := obtenerEdad(employees, nombre)
	fmt.Println("La edad de " + nombre + " es: " + strconv.Itoa(edadNombre))

	//Saber cuántos de sus empleados son mayores de 21 años.
	edadMayor := 21
	mayoresA := mayoresAEdad(employees, edadMayor)
	fmt.Println("Mayores a " + strconv.Itoa(edadMayor) + ": " + strconv.Itoa(mayoresA))
	//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	employees["Federico"] = 25
	//Eliminar a Pedro del mapa.
	delete(employees, "Pedro")
	//mostrar cómo queda el mapa:
	fmt.Println(employees)
}

func obtenerEdad(employees map[string]int, name string) int {
	return employees[name]
}

func mayoresAEdad(employees map[string]int, age int) int {
	amount := 0
	for _, value := range employees {
		if value > age {
			amount++
		}
	}
	return amount
}
func procesarMesConSwitch(mes int) string {
	var nombreMes string
	switch mes {
	case 1:
		nombreMes = "enero"
	case 2:
		nombreMes = "febrero"
	case 3:
		nombreMes = "marzo"
	case 4:
		nombreMes = "abril"
	case 5:
		nombreMes = "mayo"
	case 6:
		nombreMes = "junio"
	case 7:
		nombreMes = "julio"
	case 8:
		nombreMes = "agosto"
	case 9:
		nombreMes = "septiembre"
	case 10:
		nombreMes = "octubre"
	case 11:
		nombreMes = "noviembre"
	case 12:
		nombreMes = "diciembre"
	default:
		nombreMes = "El mes ingresado no es válido"
	}

	return nombreMes
}

func procesarMes(mes int) string {
	var meses = []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}

	if mes < 1 || mes > 12 {
		return "El mes ingresado no es válido"
	}

	return meses[mes-1]
}

func determinarPrestamo(edad int, empleado bool, antig int, sueldo int) {
	apto := true
	mensaje := "No se puede otorgar un préstamo a este cliente debido a que:\n"

	if edad <= 22 {
		apto = false
		mensaje += "\tNo es mayor de 22 años.\n"
	}

	if !empleado {
		apto = false
		mensaje += "\tNo es un empleado.\n"
	}

	if antig <= 1 {
		apto = false
		mensaje += "\tNo tiene más de un año de antiguedad.\n"
	}

	if apto {
		mensaje = "Se otorgará un préstamo al empleado"
		if sueldo > 100000 {
			mensaje += " sin interés."
		} else {
			mensaje += " con interés."
		}
	}

	fmt.Println(mensaje)
}

func procesarPalabra(palabra string) {
	fmt.Println("Palabra: " + palabra)
	cant := cantLetras(palabra)
	fmt.Println("Cantidad de letras: " + strconv.Itoa(cant))
	printLetras(palabra)
}

func cantLetras(palabra string) int {
	return len(palabra)
}

func printLetras(palabra string) {
	fmt.Print("Letras: ")
	for _, c := range palabra {
		fmt.Print(string(c) + "  ")
	}
	fmt.Println()
}
