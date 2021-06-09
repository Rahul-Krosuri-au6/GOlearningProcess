package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
var wg sync.WaitGroup
func main(){
	fmt.Println("hey")
	var counter int64
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++{
		go func ()  {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
	fmt.Println("Go Routines:", runtime.NumGoroutine())
}