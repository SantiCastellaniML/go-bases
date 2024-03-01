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
	Subject  Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e)
}

func (p Person) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DateOfBirth: %s", p.ID, p.Name, p.DateOfBirth)
}

func (e Employee) String() string {
	return fmt.Sprintf("ID: %d, Position: %s, Subject: [%s]", e.ID, e.Position, e.Subject)

}

func main() {
	p := Person{1, "Santiago", "16/08/2001"}
	e := Employee{1, "Developer", p}

	e.PrintEmployee()
}
