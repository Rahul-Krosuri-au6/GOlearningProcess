package main

import "fmt"
var y = 54
var z string
var a float32
func main() {
	x := 45
	y = x + y
	fmt.Printf("%b\t%d\t%v\t%#v\t%x\t%#x\n", y, y, y, y, y, y)
	fmt.Printf("%T\n", y)
	z = `sdsd
aaasddf
asdasd`
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	foo()
}

func foo() {
	fmt.Println("again:", y)
}
