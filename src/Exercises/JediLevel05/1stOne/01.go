package main
import "fmt"
type person struct{
	first string
	last string
	favIceCreams []string
}
func main(){
	p1 := person{
		first: "Rahul",
		last: "krosuri",
		favIceCreams: []string{"choco", "butterscotch"},
	}
	p2 := person{
		first: "Chandu",
		last: "krosuri",
		favIceCreams: []string{"vanilla", "blackcurrent"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCreams{
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCreams{
		fmt.Println(i, v)
	}
}