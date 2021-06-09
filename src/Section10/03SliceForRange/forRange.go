package main
import "fmt"
func main(){
	x := []int{4, 5, 6, 7, 8} 
	fmt.Println(x)
	for i,v  := range x{
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
	for i := 0; i < len(x); i++{
		fmt.Println(x[i])
	}
}