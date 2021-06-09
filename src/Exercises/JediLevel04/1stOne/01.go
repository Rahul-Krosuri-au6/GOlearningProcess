package main
import "fmt"

func main(){
	var x [5]int
	for i:=0; i<5; i++{
		x[i] = i
	}
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	for i, v:= range x{
		fmt.Println(i, v)
	}
}