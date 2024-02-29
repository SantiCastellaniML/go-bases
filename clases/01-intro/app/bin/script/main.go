package main

import (
	//	"fmt"
	"fmt"
	"go-bases/intro/app/lib/calculator"

	"github.com/fatih/color"
)

func main() {
	color.Red("Hello, World!")

	fmt.Println(calculator.Value)
	//color.Green(calculator.Value) //para poder utilizarlo puedo usar el nombre que tengo en el último segmento del import.
}
