package matematica

//Calculo executa operacoes recebendo o tipo e os valores
func Calculo(funcao func(int, int) int, numA int, numB int) int {
	return funcao(numA, numB)
}

//Multiplicador multiplica dois valores
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor divide dois valores
func Divisor(x int, y int) (resultado int) {
	resultado = x / y
	return
}

//DivisorComResto calcula e retorna o resultado da divisão e o resto
func DivisorComResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
