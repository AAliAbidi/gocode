package main

import "fmt"

func main() {
	fmt.Println("This is my first fmt program")
	foo()
	fmt.Println("Adding something more")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	println("This is new function by Google.")
}
func bar() {
	println("Exit")
}
