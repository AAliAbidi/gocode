package main

import "fmt"

var x bool

func main() {
	x := 10
	fmt.Println(x == x)
	fmt.Println(x != x)
	fmt.Println(x >= x)
	fmt.Println(x <= x)
	fmt.Printf("%T\n", x)
}
