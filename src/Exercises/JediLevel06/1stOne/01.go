package main
import "fmt"
func main(){
	fmt.Println(foo())
	fmt.Println(bar(1997, "born I was"))
}

func foo () int{
	return 45
}

func bar(xi int, s string) (int, string){
	return xi, s
}