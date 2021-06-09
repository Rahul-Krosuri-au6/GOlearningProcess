package main
import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
var mu sync.Mutex
func main(){
	counter := 0
	gs := 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	wg.Add(gs)
	for i:= 0; i < gs; i++{
		go func ()  {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}