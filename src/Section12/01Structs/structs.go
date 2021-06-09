package main
import "fmt"
type person struct{ 
	name string
	age int
}
func main(){
	p1 := person{
		name: "Rahul",
		age: 23,
	}
	p2 := person{
		name: "Chandu",
		age: 21,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.name, p1.age)
	fmt.Println(p2.name, p2.age)
}