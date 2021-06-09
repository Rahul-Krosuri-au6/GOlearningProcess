package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
var wg sync.WaitGroup
func main(){
	var counter int64
	gs := 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	wg.Add(gs)
	for i:= 0; i < gs; i++{
		go func ()  {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}