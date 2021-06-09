package main
import "fmt"

func main(){
	defer bar()
	defer foo()
	// if we use defer keyword before the function it asks the compiler not to run this function
	// until the other function returns the value. This is used if there are many functions and
	// if running some of them takes up more space and prevets the execution or delays it
}

func foo(){
	fmt.Println("foo here")
}

func bar(){
	fmt.Println("bar here")
}