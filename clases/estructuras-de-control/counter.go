package main

import "fmt"

func main() {
	var counter = 0
	var limit = 10

	for counter < limit {
		counter++
	}

	counter = 0
	for {
		//condición de corte:
		if counter >= limit {
			break
		}
		fmt.Println(counter)
		counter++
	}
}
