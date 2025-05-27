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

	for i := 0; i < len(pecas_carro); i++ { 
		fmt.Println(i,".",pecas_carro[i])
	}
}