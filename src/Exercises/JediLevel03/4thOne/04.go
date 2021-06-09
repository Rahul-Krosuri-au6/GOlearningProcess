package main
import "fmt"
func main(){
	alive := 1997
	for {
		if alive > 2021 {
			break
		}
		fmt.Println(alive)
		alive++
	}  
}