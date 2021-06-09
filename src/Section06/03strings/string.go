package main

import "fmt"

func main () {
	s := "hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s) // Converts into array of ascii numbers for each character
	fmt.Println(bs) 
	fmt.Printf("%T\n", bs) // type of this
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i]) // Prints the UTF-8 for the characters 
	}
	for i, v := range s {
		fmt.Printf("%d\t%c\n", i, v) // Prints the index and the characters at that particular index to show bytes use %d instead of %c
	}
}