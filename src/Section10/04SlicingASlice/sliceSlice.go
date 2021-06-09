package main
import "fmt"
func main(){
	x := []int{4, 5, 6, 7, 8} 
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:]) // starts from one goes to end
	fmt.Println(x[1:3])// starts at 1 and goes to positions 3 but doesnt print the number at position 3
}