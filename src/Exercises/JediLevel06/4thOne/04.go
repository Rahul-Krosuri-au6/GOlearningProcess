package main
import "fmt"
type person struct{
	first string
	last string
	age int
}
type human interface{
	speak()
}
func main(){
	p1 := person{
		first: "Rahul",
		last: "Krosuri",
		age: 23,
	}
	p2 := person{
		first: "Sri",
		last: "Chandan",
		age: 21,
	}
	p1.speak()
	p2.speak()
	bar(p1)
	bar(p2)
}

func (p person) speak(){
	fmt.Println("Hi!!! I am", p.first,  p.last, "and I am", p.age, "years old")
}

func bar (h human) {
	fmt.Println("I am a human and my name is", h.(person).first, h.(person).last, "and my age is", h.(person).age)
}
