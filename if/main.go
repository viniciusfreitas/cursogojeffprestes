package main

import "fmt"

func main() {

	meses := 6
	situacao := true

	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo.")
	}

	if !situacao {
		fmt.Println("Ele está em dia.")
	}

	if descricao, status := tempoDivida(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
	}

	fmt.Println("Obrigado por nos consultar.")

}

func tempoDivida(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}
	descricao = "O cliente está em dia."
	return
}
