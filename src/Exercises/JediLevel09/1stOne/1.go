package main
import ("fmt"
"runtime"
"sync"
)
var wg sync.WaitGroup
func main(){
	fmt.Println("GO Routine number 1")
	fmt.Println("GO routines: ", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	fmt.Println("GO routines: ", runtime.NumGoroutine())
	go bar()
	fmt.Println("GO routines: ", runtime.NumGoroutine())
	wg.Wait()
}

func foo(){
	fmt.Println("Go routine number 2")
	wg.Done()
}
func bar(){
	fmt.Println("Go routine number 3")
	wg.Done()
}