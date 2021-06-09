package main
import "fmt"

func main(){
	j := 42
	x := make([]int, 10, 15)
	for i:=0; i<10; i++{
		x[i] = j
		j++
	}
	fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// fmt.Printf("%T\n",x)
	// for i, v:= range x{
	// 	fmt.Println(i, v)
	// }
	x = append(x[:0],x[:5]... )
	fmt.Println(x)
}