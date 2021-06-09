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
	fmt.Println("Before changing my name is", p1.name, "and my age before changing is", p1.age)
	changeme(&p1)
	fmt.Println("After changing my name is", p1.name, "and my age after changing is", p1.age)
}

func changeme(p *person) {
	(*p).name = "Krishna Sai Rahul Krosuri" //can also be written as p.name
	p.age = 24 // can also be written as (*p).age
}