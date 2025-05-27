package main

import "fmt"

func main() { 
	pecas_carro := []string {
		"Motor",
		"Suspensão",
		"Volante",
		"Painel",
		"Escape",
		"Turbina"}

	//fatia := pecas_carro[0:(len(pecas_carro) - 1)]
	//fmt.Println(fatia)

	//excluir um item do slice, [:1] isso significa que estou partindo do 0 e chegando até o número do slice desejado, que no caso é o 1, a partir disso eu coloco outro slice[:3], mas esse delimita que meu slice peca_carro parta do 3, então pela lógica ele tira o espaço 2 que é a "Suspensão", e assim que ficará = 0, 1, 3, 4 e 5.

	pecas_carro = append(pecas_carro[:1], pecas_carro[2:]...)
	for i:= 0; i < len(pecas_carro); i++ { 
		fmt.Printf("%v. %v\n", i, pecas_carro[i])
	}
}