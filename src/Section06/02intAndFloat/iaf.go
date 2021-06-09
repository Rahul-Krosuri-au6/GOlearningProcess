package main

import (
	"fmt";	"runtime"
)


var x int
var y float64

func main(){
	x = 55
	y = 56.2564
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(runtime.GOOS) //operating system on which go runs i.e your operating system
	fmt.Println(runtime.GOARCH) //architecture on which go runs
}