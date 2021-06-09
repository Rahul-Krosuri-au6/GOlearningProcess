package main
import "fmt"
func main(){
	foo()
	var name string
	var firstname string
	var lastname string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your first name:")
	fmt.Scanln(&firstname)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lastname)
	hello(name)
	fmt.Println(returnFunc("hey guys"))
	fmt.Println(twoP(firstname, lastname))
}
// basic func
func foo(){
	fmt.Println("Hello from another function")
}
// func with a parameter
func hello(s string){
	fmt.Println("hello", s)
}
// func which returns a string or anything we like
func returnFunc(st string) string{
	return fmt.Sprint("This is a return function, ", st )
}
// func which takes two parameters and returns them
func twoP(fn, ln string) string{
	return fmt.Sprint(fn, " ", ln, ` says "hello"`)
}
// parameters are passed when we write the function and argumesnts are passed when the function
// is called