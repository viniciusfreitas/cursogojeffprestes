package main

import (
	"fmt"

	"github.com/cursogojeffprestes/calculadora/matematica"
)

func main() {

	resultado := matematica.Calculo(matematica.Multiplicador, 2, 3)
	fmt.Printf("2 x 3 = %d\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 10, 2)
	fmt.Printf("10 ÷ 2 = %d\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("12 ÷ 5 = %d | resto = %d\n", resultado, resto)
}
