//  Number of Methods attactched to a set is called method sets
package main
import "fmt"
const PI = 3.14
type square struct{
	side float64
}
type circle struct{
	radius float64
}
type shape interface{
	area() float64
}

func (s *square) area() float64{ 
	return s.side * s.side
}
func (c *circle) area() float64{
	return c.radius * c.radius * PI
}

func info(sh shape){
	// switch sh.(type){
	// case square:
	// 	fmt.Println("The area of square is", sh.(square).area())
	// case circle:
	// 	fmt.Println("The area of circle is", sh.(circle).area())
	// }
	fmt.Println("Area is", sh.area())
}

func main(){
	fmt.Println("hey")
	s1 := square{
		side: 12,
	}
	c1:= circle{
		radius: 12,
	}
	// info(&s1) // can recieve both pointer and non pointer values i.e runs for &s1 too when no swtich case is present
	// info(&c1)
	info(&s1)
	info(&c1) // if the parameter type is a pointer that argument must also be a pointer value
}

// Rules:-
// No pointer parameters can take both pointer and non pointer values as arguments
// Pointer parameters can take only pointer values
