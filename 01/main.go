package main

import "fmt"

func main() {
	fmt.Println("La suma de 2 + 4 es:", miSuma(2, 4))
	fmt.Println("La suma de 1 + 5 es:", miSuma(1, 5))
	fmt.Println("La suma de 6 + 7 es:", miSuma(6, 7))

}

func miSuma(xi ...int) int { //recibe cantidad variable de enteros
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum + 1 //el +1 es para que nos de error y probar el error
}
