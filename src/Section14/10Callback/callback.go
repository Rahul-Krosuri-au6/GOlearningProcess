package main
import "fmt"
func main () {
	ii := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(sum(ii...))// Prints the total of number in ii
	s2 := even(sum, ii...) // when sum is used as an callback the slice is passed as an argument to sum and as given in the function on ly even numbers are printed
	fmt.Println(s2)
	s3 := odd(sum, ii...)
	fmt.Println(s3)
}

func sum(x ...int) int{
	total := 0
	for _,v := range x{
		total += v
	}
	return total
}

func even(f func(y ...int) int, yi ...int) int{ // takes in a function which takes variadic int parameter and a variadic int parameter and asiigns it tothe call back function in the parameter
	y := []int{} //empty slice is created
	for _,v := range yi{
		if (v % 2 == 0){
			y = append(y, v) //here even numbers are appended to the list
		}
	}
	return f(y...) //the above slice is unfurled and passed as an argument into the function
}

func odd(f func(z ...int) int, zi ...int) int{
	z := []int{}
	for _,v := range zi{
		if v % 2 != 0 {
			z = append(z, v)
		}
	}
	return f(z...)
}