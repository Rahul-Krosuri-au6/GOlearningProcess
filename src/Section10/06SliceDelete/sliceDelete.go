package main
import "fmt"

func main(){
	x := []int{1, 2, 3, 4, 5} 
	fmt.Println(x)
	x = append(x, 6, 7, 8)
	fmt.Println(x)
	y := []int{9, 10, 11, 12}
	x = append(x, y...) // ... after y implies append every element of y to x if before y it implies that it takes any number of data of the types of y
	fmt.Println(x)
	x = append(x[:2], x[4:]...) // removes elements from index 2 and 3
	fmt.Println(x)
}