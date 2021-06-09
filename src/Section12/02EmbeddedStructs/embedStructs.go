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
	fmt.Println(sa2)
	fmt.Println(sa1.name, sa1.age, sa1.ltk) // if there is a name collision within the embedded structs we can use sa1.person.name
	fmt.Println(sa2.name, sa1.age, sa2.ltk)
	// instead of declaring person struct above the main function we can declare the struct body
	// in the p1 itself i.d p1 := struct{
									// name string
									// age int
								// }{
										// 	name: "James Bond",
										// 	age: 35,
									// },
	// as shown above we can declare the struct in that way and struct is called annd struct
	// can be called anonymus struct as it doesn't have any name. This method is used when we need a
	// struct only for one or two variables in a big code which prevents code pollution and
	// struct clutter
}