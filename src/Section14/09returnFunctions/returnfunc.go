package main
import "fmt"
func main(){
	// s := foo()
	fmt.Println(foo()) // can be written as foo()
	// x := bar()
	fmt.Println(bar()()) // () to call the anonymous function else gives the address (can also be written as bar()())
	fmt.Printf("%T\n", bar()) // type of x is func() int but type of x() is int
}

func foo() string{
	return "Hi from foo"
}

func bar() func() int{
	return func() int{
		return 45
	}
}