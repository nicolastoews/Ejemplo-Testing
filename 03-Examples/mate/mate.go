//Package mate nos ayuda a comprobar que sabes sumar
package mate

//sum suma cantidad ilimitada de ints
func sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
