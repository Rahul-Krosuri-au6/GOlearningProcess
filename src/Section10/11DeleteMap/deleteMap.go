package main
import "fmt"
func main(){
	m := map[string]int{
		"Rahul": 23,
		"chandu": 21,
	}
	fmt.Println(m)
	m["Madhavi"] = 48 // adds a new key and value to the given dictionary
	for k, v := range m{
		fmt.Println(k, v)
	}
	delete(m, "chandu") //delete a key "chandu" from map m
	fmt.Println(m)
	// we can also delete something that doesn't exist no error is given
	delete(m, "Sri")
	fmt.Println(m)
}