package main
import "fmt"
func main(){
	switch true{
	case true:
		fmt.Println("This should be printed1")
		fallthrough
	case false:
		fmt.Println("This should not be printed")
	}
}