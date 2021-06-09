package main
import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
	counter := 0
	gs := 100
	// x := 0
	wg.Add(gs)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	for i:= 0; i < gs; i++{
		go func ()  {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		// go func ()  {
		// 	x = counter
		// 	x++
		// 	counter = x
		// 	// wg.Done()
		// }()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
	// fmt.Println("Counter: ", counter)
}