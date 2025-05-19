package main

import "fmt"

func main() {

	/*for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}*/

	i := 1

	for {
		if i <= 10 {
			fmt.Println("i é menor que dez", i)
			i++
		} else {
			fmt.Println("i é maior que dez!")
			break
		}
	}
	fmt.Println("O loop foi brekado!")
	

}
