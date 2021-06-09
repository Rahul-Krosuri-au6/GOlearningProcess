package main

import "fmt"
type some int
var x some
func main(){
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
