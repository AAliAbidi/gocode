package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello", 22, true)
	fmt.Println(n)
	fmt.Println(err)
	div()
}
func div() {
	n,  := fmt.Println(2)
	fmt.Println(n)
}
