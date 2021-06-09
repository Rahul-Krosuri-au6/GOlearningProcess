package main
import "fmt"
// slice groups same items together
//  and arrays groups same items together too
func main(){
	// composite literal i.e creates a new instance everytime it is evaluated
	// composite literals can be used for slices, arrays, structs and maps
	x := []int{4, 5, 6, 7, 8} // type and values
	fmt.Println(x)
}