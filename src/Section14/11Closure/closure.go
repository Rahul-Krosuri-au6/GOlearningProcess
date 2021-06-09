//  in golang we try to keep the scope as narrow as possible and this is done by a process called closure
// closure is like a code block enclosing some variables and narrowing their scope
package main
import "fmt"
func main(){
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // increments the value of x as long as it is called by a
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // increments the value of x from start as we start from different memory locations
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int // this is called closure as we limit the scope of x to this function only
	return func() int{
		x++
		return x
	}
}