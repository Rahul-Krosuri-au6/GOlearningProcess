package main
import ("fmt"
"runtime"
"sync"
)
var wg sync.WaitGroup
var mu sync.Mutex
func main(){
	fmt.Println("hey")
	counter := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++{
		go func ()  {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("counter:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}