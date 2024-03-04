package internal

func operation(v1, v2 int) (sum int, sub int, mul int, div int) {
	if v2 != 0 {
		div = v1 / v2
	}

	sum = v1 + v2
	sub = v1 - v2
	mul = v1 * v2

	return //no hace falta indicar qué retorna, ya sabe qué variables debería retornar. IGUALMENTE SE PUEDE RETORNAR UN VALOR CONCRETO.
}
