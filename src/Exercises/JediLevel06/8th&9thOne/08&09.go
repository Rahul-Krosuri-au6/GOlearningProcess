package main
import "fmt"
func main(){
	xi := []int{1,2,3,4,5,6,7}
	fmt.Println(sum(xi...))
	fmt.Println(add(sum, xi...))
	fmt.Println(foo()())
}

func add(f func(s ...int) int, yi ...int) int{
	y:= []int{}
	for _,v := range yi{
		if v % 2 == 0{
			y = append(y, v)
		} 
	}
	return f(y...)
}

func sum(x ...int) int{
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}

func foo() func() string{
	return func() string{
		return "Func which return a func"
	}
}