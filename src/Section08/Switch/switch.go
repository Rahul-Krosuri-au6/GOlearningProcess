package main
import "fmt"
func main(){
	switch 5{ // 5 is like a expression i.e a true statement that needs to be achieved can also be any data type (string, bool etc)
	case 1 :
		fmt.Println("not 5")
		fallthrough // allows code to move forward even if true is achieved at a particular step
	case 2 :
		fmt.Println("not 5")
		fallthrough
	case 3 :
		fmt.Println("not 5")
		fallthrough
	case 4 :
		fmt.Println("not 5")
		fallthrough
	case 5 :
		fmt.Println("it is 5")
		fallthrough
	default:
		fmt.Println("This is the end case or default case") // optional but good to have in case all fails
	}
}