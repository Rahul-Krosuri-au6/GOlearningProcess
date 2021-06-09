package main

import "fmt"

func main(){
	fmt.Println("Hello Golang!!")
	foo()
	for i:=0; i < 10; i++ {
		if i % 2 == 0{
			fmt.Println(i)
		}
	}
}

func foo(){
	fmt.Println("This is func!")
}