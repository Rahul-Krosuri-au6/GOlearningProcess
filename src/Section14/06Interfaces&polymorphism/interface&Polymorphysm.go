package main
import "fmt"
type person struct{
	name string
	age int
}
type secretagent struct{
	person
	ltk bool
}

type human interface{
	speak() // an interface called human is created and every struct that has the method speak attatched
	// to it is also of the type human i.e the interface allows the value to have two types
	// if no method is given then every type with no method becomes the type of this interface
}
// conversion is converting one type to another
type hotdog int
var x hotdog
var y int
func main(){
	x = 55
	y = 45
	y = int(x) // converting the type hotdog to int
	p1:= person{
		name: "rahul",
		age: 23,
	}
	sa1 := secretagent{
		person: person{
			name: "James Bond",
			age: 35,
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			name: "Miss moneypenny",
			age: 30,
		},
		ltk: false,
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
	p1.speak()
	bar(sa1) // the bar function takes in arguments which are only of type human as sa1, sa2 and p1 are
	bar(sa2) //all are of types human it allows them as parameters. This is also called polymorphism as 
	bar(p1) // as the same function is used for various types and different output is given
	fmt.Printf("%T\n", sa1)
	fmt.Printf("%T\n", sa2)
	fmt.Printf("%T\n", p1)
}
// method arch is func (s struct) funcName({parameters}) {returns} {...code...}
// method takes exactly one reciever and doesn't need parameters
func (s secretagent) speak(){
	fmt.Println("I am ", s.name, " I am ", s.age, " years old, and I have license to kill which is ", s.ltk)
}
func (p person) speak(){
	fmt.Println("I am a person with name ", p.name, " I am ", p.age, " years old")
}
// if we need to use custom messgaes for type it has then we use assertion
func bar (h human) {
	switch h.(type) { // this tells if it to accept if only of same type
	case secretagent:
		fmt.Print("I am a human called ", h.(secretagent).name, " and I am a secret agent\n") // the type name in () is called assertion as it tells h that it is a struct of secret agent
	case person:
		fmt.Print("I am a human called ", h.(person).name, " and I am a normal person\n")
	}
	
}