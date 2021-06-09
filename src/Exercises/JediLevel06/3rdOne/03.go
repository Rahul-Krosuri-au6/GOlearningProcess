package main
import "fmt"
func main(){
	defer foo()
	fmt.Println("foo was defered")
	bar()
}

func foo(){
	defer func ()  {
		fmt.Println("I am in a defered func in foo and will be printed at the last line")
	}()
	fmt.Println("I will be printed at the last")
}

func bar(){
	fmt.Println("I will be printed before foo even though declared later because of defer")
}

