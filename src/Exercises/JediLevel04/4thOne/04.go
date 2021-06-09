package main
import "fmt"

func main(){
	j := 42
	x := make([]int, 10, 15)
	for i:=0; i<10; i++{
		x[i] = j
		j++
	}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y:= []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}