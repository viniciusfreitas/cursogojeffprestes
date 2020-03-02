package main

import "fmt"

//Imovel é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("Conteúdo da variável original: %+v\r\n", casa)
	fmt.Printf("Endereço de memória da variável original: %p\r\n", casa)

	chacara := Imovel{17, 38, "Chacara", 2800000}
	fmt.Printf("Imóvel: %p | %v\r\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Imóvel: %p | %v\r\n", &chacara, chacara)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
