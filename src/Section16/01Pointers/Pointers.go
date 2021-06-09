package main
import "fmt" 
func main(){
	x := 45
	y := 55
	fmt.Println("Before swap", x, y)
	swap(&x, &y)
	fmt.Println("After swap", x, y)
}
func swap (x *int, y *int){ // *int tells us that it is a pointer
	temp := *x //*x tells that it is a type of operator which gives the address
	fmt.Println("Value of 1st variable before swapping", *x)//gives the value of x i.e dereferences the pointer i.e follows the address and gives the value stored at that address
	fmt.Println("Address of 1st variable before swapping", &x)//gives the address of x
	fmt.Println("Value of 2nd variable before swapping", *y)
	fmt.Println("Address of 2nd variable before swapping", &y)
	*x = *y
	*y = temp
	fmt.Println("Value of 1st variable after swapping", *x)
	fmt.Println("Address of 1st variable after swapping", &x)
	fmt.Println("Value of 1st variable after swapping", *y)
	fmt.Println("Address of 1st variable after swapping", &y)

}
// if we are asked to find the value stored at an address we use "*&address"