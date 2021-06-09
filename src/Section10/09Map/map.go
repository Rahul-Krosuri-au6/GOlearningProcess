package main
import "fmt"
func main(){
	m := map[string]int{
		"Rahul": 23,
		"chandu": 21,
	}
	fmt.Println(m)
	fmt.Println(m["Rahul"])
	fmt.Println(m["Madhavi"]) // not present in the original dict so zero value
	fmt.Printf("%T\n", m)
	v, ok := m["madhavi"] // ok to check if it is present in dict gives true or false
	// v1, ok1 := m["Rahul"] //gives ok as true and executes the if statement
	fmt.Println(v) // gives value of the key you want to find in a dict
	fmt.Println(ok) // prints true if present and false if not, In thi case false as the key is not present in the dict
	// fmt.Println(v1)
	// fmt.Println(ok1)
	if ok {// ok1 
		println("this is in the dictionary") // print this if present i.e ok is true
	}
	// v, ok can be any names but we this as a concention
}
