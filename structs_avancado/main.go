package main

import (
	"encoding/json"
	"fmt"

	"github.com/cursogojeffprestes/structs_avancado/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor = 60000

	fmt.Printf("O imóvel é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Printf("A casa em JSON: ", string(objJSON))

}
