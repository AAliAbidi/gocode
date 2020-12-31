package main

import "fmt"

var y int = 10

//Creating your own variable type
type sun int

var b sun = 1000

func main() {
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%T\n", b)
}
