package main
import "fmt"
func main(){
	x:=25 //change x value to 23 and 22 to cjeck if other statements are working
	if x%5==0{
		fmt.Println("divisible by 5")
	}else if x % 11 ==0{
		fmt.Println("Divisible by 11")
	}else{
		fmt.Println("NOt divisible by any")
	}
}