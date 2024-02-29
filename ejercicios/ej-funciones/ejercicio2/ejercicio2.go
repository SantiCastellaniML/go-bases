package main

func main() {

}

func calcularPromedio(numeros ...float32) (float32, string) {
	var suma float32
	for _, numero := range numeros {
		if numero < 0 {
			return 0, "No se puede indicar una nota negativa"
		}
		suma += numero
	}

	return suma / float32(len(numeros)), ""
}
