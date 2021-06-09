package main

import "fmt"

func main(){
	fmt.Println("hey")
	foo()
	// This is an anonymus function it has no name except for keyword func
	func(){
		fmt.Println("This is an anonymous func with no  parameters")
	}() // this is needed to call the function just like calling the function foo()
	// with parameters
	func(x int){
		fmt.Println("This is an anonymous function with an int parameter ", x)
	}(23)
	// with return statements
	x :=func (x int) int{
		return x
	}(55)
	fmt.Println("the reutrn from the anonymous function is ", x)

}

//  This is a normal function called foo
func foo(){
	fmt.Println("This is from foo")
}
