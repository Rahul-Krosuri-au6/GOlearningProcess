package main
import "fmt"
func main(){
	x := []int{1, 2, 3, 4, 5, 6, 7,}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(xi ...int) int{
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}

func bar(ii []int) int{
	sum := 0
	for _, v := range ii{
		sum += v
	} 
	return sum
}