package model

//Imovel é a struct para imóveis
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor define o valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
