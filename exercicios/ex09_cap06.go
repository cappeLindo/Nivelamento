package main 

import "fmt"

func main() { 
	
	esporteFav := "beachtenis"
	switch esporteFav{
	case "voleiball":
		fmt.Println("massa, voleiball é legalzin")
	case "futebol": 
		fmt.Println("Ai ss, melhor esporte que tem")
	case "beachtenis": 
		fmt.Println("🏳️‍🌈")
	case "automobilismo":
		fmt.Println("Sabe muito")
	default:
		fmt.Println("Não temos esse esporte catalogado")
	}
}