package main
import "fmt"

func main(){
	fmt.Println("hey")
	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum of all digits is: ", x)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x{
		sum += v
	}
	return sum
}