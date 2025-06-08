package main

import "fmt"

func main() {

	s := struct {
		nome 	string
		idade 	int8
		carac	[]string
		hoobs	map[int]string
	} {
		nome: "vitor",
		idade: 17,
		carac: []string{"loiro", "olho verde", "1.76m"},
		hoobs: map[int]string{
			1 : "codar, acho que é um dos poucos hoobs que é benefico",
			2 : "tenho hoob em jogo de tiro",
		},
	}
	fmt.Printf("Nome: %v\nIdade: %v\nCaracteristicas: %v\n", s.nome, s.idade, s.carac)
	for i, v := range s.hoobs { 
		fmt.Printf("hoob %v : %v\n", i, v)
	}
}

