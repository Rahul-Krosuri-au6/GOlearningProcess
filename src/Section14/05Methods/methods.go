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
func main(){
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
}
// method arch is func (s struct) funcName({parameters}) {returns} {...code...}
// method takes exactly one reciever and doesn't need parameters
func (s secretagent) speak(){
	fmt.Println("I am ", s.name, " I am ", s.age, " years old, and I have license to kill which is ", s.ltk)
}