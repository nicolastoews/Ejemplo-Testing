//Cumple con el requisito de que el archivo tiene que terminar con _test.go
//También cumple con que esté en el mismo package de lo que estamos testeando
//Nota personal: Este ejemplo está en otro lado por que la ruta del archivo de pruebas no puede contener espacios.

package main

import "testing"

func TestMiSuma(t *testing.T) { //No es necesario que se llame igual a la función que vamos a testear pero es una buena práctica
	v := miSuma(2, 8)
	if v != 10 {
		t.Error("Expected", 10, "got", miSuma(2, 8))
	}
}
