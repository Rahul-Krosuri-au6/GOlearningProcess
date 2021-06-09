package main
import "fmt"

func main(){
	y := forFun()
	x := foo(y...)// ... to tell that the slice of numbers is to be taken as it is instead of creating a new slice
	// we can also pass no parameter as variadic means we can pass as many as ints as possible
	// which also includes zero
	fmt.Println("The sum of all digits is: ", x)
}

func foo(x ...int) int {  //... must be the final parameter any thing after that is not accepted
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x{
		sum += v
	}
	return sum
}
func forFun() []int{
	var length int
	var capacity int
	fmt.Println("Enter the lenght: ")
	fmt.Scanln(&length)
	fmt.Println("Enter the capacity: ")
	fmt.Scanln(&capacity)
	xi := make([]int, length, capacity)
	var num int
	for i:=0; i < length; i++{
		fmt.Println("Enter a number to the slice: ")
		fmt.Scanln(&num)
		xi[i] = num
	}
	return xi
}