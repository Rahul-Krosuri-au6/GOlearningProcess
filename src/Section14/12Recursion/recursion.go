package main
import "fmt"

func main(){
	fmt.Println("This is from recursion ", factorial(4))
	fmt.Println("This is from loop ", factUsingLoop(4))
}

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}
// using loop as a part of challenge
func factUsingLoop(n int) int{
	total := 1
	for ;n > 0; n--{
		total *= n
	}
	return total
}