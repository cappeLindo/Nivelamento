package main

import "fmt"


func main() {
	estados := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	for indice, valor := range estados { 
		fmt.Printf("%v - %v\n", indice + 1, valor)
	}
	
	for i := 1; i < len(estados); i++ { 
		fmt.Printf("existem %v estados no Brasil \n", i)
		fmt.Println(estados[i])
	}
}