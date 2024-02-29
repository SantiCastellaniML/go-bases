package ejercicio3

const (
	SalaryA = 3000.0
	SalaryB = 1500.0
	SalaryC = 1000.0
)

func main() {

}

func CalcularSalario(minutos int, categoria string) (salario float32, err string) {
	var minSalary float32
	switch categoria {
	case "C":
		minSalary = SalaryC / 60
		return minSalary * float32(minutos), ""
	case "B":
		minSalary = SalaryB / 60
		salary := minSalary * float32(minutos)
		salary *= 1.2
		return salary, ""
	case "A":
		minSalary = SalaryA / 60
		salary := minSalary * float32(minutos)
		salary *= 1.5
		return salary, ""
	default:
		return 0, "Categoria no valida"
	}
}
