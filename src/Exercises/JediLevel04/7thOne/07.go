package main
import "fmt"

func main(){
	x := []string{"james", "bond", "shaken but not stirred"} 
	y := []string{"Miss", "MoneyPenny", "Hello James"}
	for i,v:= range x{
		fmt.Println(i, v)
	} 
	for i,v:= range y{
		fmt.Println(i, v)
	}
	xy := [][]string{x, y}
	fmt.Println(xy)
	for i,v:= range xy{
		fmt.Println(i, v)
		for j, val:= range v{
			fmt.Println(j, val)
		}
	} 
}