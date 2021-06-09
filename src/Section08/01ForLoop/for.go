package main

import "fmt"

func main(){
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Println(i, j)	
// 	}
// }
// 	// for loop with a single statement is replacement for while in golang
// 	a := 55
// 	b := 5
// 	for a > b {
// 		z := a/5
// 		fmt.Println(z)
// 		a = z
// 		b += 1
// 	}
	// break and continue statements
	// x := 100
	// for i := 1; i <= x; i++{
	// 	if (x == i){
	// 		continue
	// 	}
	// 	fmt.Println(x/i)
	// 	x--
	// }
	// Even numbers
	// x := 0
	// for {
	// 	x++
	// 	if (x > 50){
	// 		break
	// 	}
	// 	if (x%2 != 0){
	// 		continue
	// 	}
	// 	fmt.Println(x)
	// }
	// Numbers to ascii letters
	x := 0
	for {
		if (x > 122){
			break
		}
		fmt.Printf("%c", x)
		x++
	}
}