package main

import "fmt"

var x  = 42
var y string = "James Bonde"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(s)

}

