package main
import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func main(){
	fmt.Println("hey")
	counter := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++{
		go func ()  {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("counter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}