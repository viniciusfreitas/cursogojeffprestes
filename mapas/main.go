package main

import (
	"fmt"

	"github.com/cursogojeffprestes/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa amarela"] = casa

	fmt.Println("há uma casa amarela no meu cache?")
	imovel, achou := cache["Casa amarela"]
	if achou {
		fmt.Printf("Sim, achei a casa: %v\r\n", imovel)
		fmt.Printf("Sim, achei a casa: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apartamento novo"
	apto.X = 29
	apto.Y = 98
	apto.SetValor(50000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens existem no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa amarela"]
	if achou {
		delete(cache, imovel.Nome)
		fmt.Println(imovel.Nome, "apagada")
	}

	fmt.Println("Quantos itens existem agora? ", len(cache))

}
