package main
import "fmt"

func main(){
	x := make([]int, 10, 11) // make (type, len, capacity)
	x[0] = 10
	x[9] = 15
	// x[10] = 12 gives errors because index out of range
	x = append(x, 12) // doesnot give error as the lenght is below capacity
	fmt.Println(x)
	fmt.Println(len(x))
	// Capacity is the underlying elements that we can add to the array above its length
	fmt.Println(cap(x))
	x = append(x, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//  when appending goes beyond the capacity a new array with twice the capacity is created and all
	// the elements from previous array are copied into this and adiitional size equal to the length
	// of the previous array is present
}