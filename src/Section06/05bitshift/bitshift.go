// bitshifting is shifting the bits left or right we use >> to shift bits to right and << to shift
// bits to left
package main

import "fmt"
// iota for bitshifting
const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main(){
	x := 5 // 101 in binary
	fmt.Println(x)
	y := x << 1 // 10 as every bit is shifted to one position to left i.e 1010
	fmt.Println(y)
	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
}