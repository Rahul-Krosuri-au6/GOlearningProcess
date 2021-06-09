package main

import "fmt"

func main(){
	x := 10
	fmt.Println(x)
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Println(y)
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}