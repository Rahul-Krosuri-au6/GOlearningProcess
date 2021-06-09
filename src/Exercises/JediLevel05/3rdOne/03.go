package main
import "fmt"
type vehicle struct{
	doors int
	color string
}

type sedan struct{
	vehicle
	luxury bool
}
type truck struct{
	vehicle
	fourWheel bool
}

func main(){
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(truck1.doors)
	fmt.Println(truck1.color)
	fmt.Println(truck1.fourWheel)
	fmt.Println(sedan1.doors)
	fmt.Println(sedan1.color)
	fmt.Println(sedan1.luxury)
}