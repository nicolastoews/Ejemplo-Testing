//Cumple con el requisito de que el archivo tiene que terminar con _test.go
//También cumple con que esté en el mismo package de lo que estamos testeando
//Nota personal: Este ejemplo está en otro lado por que la ruta del archivo de pruebas no puede contener espacios.

package main

import "testing"

func TestMiSuma(t *testing.T) { //No es necesario que se llame igual a la función que vamos a testear pero es una buena práctica
	type prueba struct {
		datos     []int //datos para cada prueba
		respuesta int   //respuesta para cada prueba
	}
	pruebas := []prueba{ //creamos varios valores del tipo prueba
		prueba{[]int{2, 4, 6}, 12},
		prueba{[]int{1, 5, 2}, 8},
		prueba{[]int{0, -1, 1}, 0},
		prueba{[]int{-10, 4, 20}, 14},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...) //en x guardamos los valores que devuelve el range de pruebas y como son varios se usa el ...
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "got", v)
		}
	}
}
