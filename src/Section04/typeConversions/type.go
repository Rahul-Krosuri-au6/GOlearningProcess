package main
import "fmt"
import "strconv"
var a int
type hotdog int
var b hotdog
var z string
func main() {
	a = 55
	b = 54
	z = "56"
	i, err := strconv.Atoi(z)
	i = i + int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println(err)
}
