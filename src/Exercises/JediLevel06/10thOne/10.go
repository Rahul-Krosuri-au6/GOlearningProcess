package main
import "fmt"
func main(){
	x := 45
	fmt.Println("hey")
	fmt.Println(foo(15))
	fmt.Println(&x)
}

func foo(s int) int{
	for i := 0; i < 10; i++{
		s = s + i
	}
	return s
}