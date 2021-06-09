package main
import "fmt"
type person struct{
	name string
	age int
}
type human interface{
	speak()
}
func main(){
	fmt.Println("hey")
	p1 := person{
		name: "Rahul",
		age: 23,
	}
	saySomething(&p1)
	p1.speak()
}

func(p *person) speak() {
	fmt.Println("name:", p.name, "age:", p.age)
}

func saySomething(say human){
	say.speak()
}