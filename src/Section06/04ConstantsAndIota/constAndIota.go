package main

import "fmt"
const a int = 55
const b = 45.25
const c = "Rahul"
const (
	d = iota // increments the value of next consts with iota refreshes to zero as when next group of iota constants are declared
	e = iota //only one iota to the first constant is enough to tell that the group is iota
	f = iota
)
func main () {
	fmt.Println(a, b, c)
	fmt.Printf("%T\t%T\t%T\n", a, b, c)
	fmt.Println(d, e, f)
	fmt.Printf("%T\t%T\t%T\n", d, e, f)
}