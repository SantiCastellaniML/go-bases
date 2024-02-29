package main

func main() {

}

func calcularSalario(minutos int, categoria string) (salario float32, err string) {
	var minSalary float32
	switch categoria {
	case "C":
		minSalary = 1000.0 / 60
		return minSalary * float32(minutos), ""
	case "B":
		minSalary = 1500.0 / 60
		salary := minSalary * float32(minutos)
		salary *= 1.2
		return salary, ""
	case "A":
		minSalary = 3000.0 / 60
		salary := minSalary * float32(minutos)
		salary *= 1.5
		return salary, ""
	default:
		return 0, "Categoria no valida"
	}
}
