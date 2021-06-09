package main
import "fmt"

func main(){
	x := make([]int, 10, 15)
	for i:=0; i<10; i++{
		x[i] = i
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n",x)
	for i, v:= range x{
		fmt.Println(i, v)
	}
}