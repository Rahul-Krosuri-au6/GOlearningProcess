package main
import "fmt"

func main(){
	x := []int{1, 2, 3, 4}
	y:= []int{5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(y)
	xy := [][]int{x, y} // creating multidimensional array
	fmt.Println(xy)
	fmt.Printf("%T\n", xy)
	fmt.Printf("%T\n", x)
}