package main

import "fmt"

func main() {
	truck1 := struct {
		doors     int
		color     string
		fourWheel bool
	}{
		doors:     2,
		color:     "white",
		fourWheel: true,
	}
	fmt.Println(truck1.color)
	fmt.Println(truck1.doors)
	fmt.Println(truck1.fourWheel)
}