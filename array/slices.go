package main

import "fmt"

func main() { 
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
    slice[2] = 555
	slice2 := append(slice, 9)
	fmt.Println(slice2)
	fmt.Println(slice[2])
}