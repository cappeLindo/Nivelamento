package main

import "fmt"

func main() { 
	pessoa := map[string][]string { 
		"vitor" : []string { 
			"gostar de jogar fps", "gosta de ser feliz", "gosta da natureza e suas paisagens",
		},
		"cr7" : []string { 
			"eu sou o melhor", "eu gosto de jogar um bola", "meu hoob é ganhar bola de ouro",
		},
		"vasco" : []string {"meu hoob é cair para série B", "sou um time de futebol", "sou o segundo time mais ntigo do Brasil",
		},
	}
	pessoa["barbeiro"] = []string {"eu gosto de cortar cabelo", "sou o melhor no degradê", "eu gosto de jogar uma pelada"}
	delete(pessoa, "vasco")
	for indice, _ := range pessoa { 
		fmt.Printf("%v\n", indice)
		for i, valor := range pessoa[indice] { 
			fmt.Printf("\t%v - %v\n", i + 1, valor)
		}
	}
}