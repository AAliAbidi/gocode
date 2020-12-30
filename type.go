package main

import "fmt"

var y = 42
var z = "Samasung Galaxy"
var x = `this is also string's`
var k = `"This is multiple", 
string`

var l string

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(k)
	fmt.Printf("%T\n", k)
	l = "james bond"
	fmt.Println(l)
	fmt.Printf("%T\n", l)
	//z = 100
	//fmt.Println(z)
}
