package main
import "fmt"
func main(){
	x :=func (x int) int{
		return x
	}(55)
	fmt.Println("The function expression assined to x has a value of ", x)
}