package main

import "fmt"

type Alumno struct {
	Name    string
	Surname string
	DNI     int
	Date    string
}

func main() {
	alumnos := []Alumno{
		{"Santiago", "Castellani", 43316372, "22/07/2019"},
		{"Ricardo", "Gonzalez", 40023072, "10/03/2017"},
	}

	for _, a := range alumnos {
		a.showAlumno()
	}
}

func (a Alumno) showAlumno() {
	fmt.Print("Name: ", a.Name, "\nApellido: ", a.Surname, "\nDNI: ", a.DNI, "\nDate: ", a.Date, "\n\n	")
}
