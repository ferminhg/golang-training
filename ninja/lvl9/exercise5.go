package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("wopwop")
	fmt.Println("numCPus:", runtime.NumCPU())

	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("numgoroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("numgoroutine:", runtime.NumGoroutine())
	fmt.Println("final value: ", counter)
}
