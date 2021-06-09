package main

import "fmt"

func main(){
	x := 54
	y := 9
	if x == y {
		fmt.Println("equal")
	}
	if x <= y {
		fmt.Println("x less than or equal to y")
	}
	if x >= y {
		fmt.Println("x greater than or equal to y")
	}
	if x != y {
		fmt.Println("x is not equal to y")
	}
	if x < y {
		fmt.Println("x is strictly less than or equal to y")
	}
	if x > y {
		fmt.Println("x strictly greater than or equal to y")
	}
}