package main

import (
	"encoding/json";	"fmt"
)

type person struct{
	Name string
	Age int
	Res []string
}
func main(){
	p1 := person{
		Name: "Rahul", 
		Age: 23,
		Res: []string{"Nandyal", "Bangalore"},
	}
	p, err := json.Marshal(p1) // convert struct to a json
	// unmarshal is making json into string it takes the json to be converted and an addressas it accepts
	// only pointers as a second argument
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println(p) // gives ascii characters as a part of output
	fmt.Println(string(p))
} 