package main
import "fmt"
func main(){
	x:="favSport"
	switch x{
	case "cricket":
		fmt.Println("Not fav Sport")
	case "Badminton":
		fmt.Println("Not fav sport")
	case "favSport":
		fmt.Println("Literal fav sport")
	}
}