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

	casa := Imovel{}
	fmt.Printf("A casa é %+v\r\n", casa)

	apartamento := Imovel{30, 19, "Meu apartamento", 200000}
	fmt.Printf("O imóvel é %+v\r\n", apartamento)

	chacara := Imovel{
		X:     29,
		Y:     39,
		Nome:  "Chacara",
		valor: 376347,
	}
	fmt.Printf("O imóvel é %+v\r\n", chacara)

	casa.Nome = "Casinha"
	casa.valor = 328476238
	casa.X = 341
	casa.Y = 126
	fmt.Printf("A casa é %+v\r\n", casa)
}
