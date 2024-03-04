package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e)
}

func (p Person) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DateOfBirth: %s", p.ID, p.Name, p.DateOfBirth)
}

func (e Employee) String() string {
	return fmt.Sprintf("ID: %d, Position: %s, Personal data: [%d, %s, %s]", e.ID, e.Position, e.Person.ID, e.Name, e.DateOfBirth)
}

func main() {
	p := Person{1, "Santiago", "16/08/2001"}
	e := Employee{1500, "Developer", p} //no hay forma de distinguir este ID del otro, por eso se pone e.Person.ID para mostrar ese.

	e.PrintEmployee()
}
