package main
import "fmt"

func main(){
	j := 42
	x := make([]int, 10, 15)
	for i:=0; i<10; i++{
		x[i] = j
		j++
	}
	x = append(x[0:3],x[6:]... )
	fmt.Println(x)
}