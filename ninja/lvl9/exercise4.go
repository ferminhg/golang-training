package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("wopwop")
	fmt.Println("numCPus:", runtime.NumCPU())

	counter := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			mutex.Lock()
			value := counter
			value++
			counter = value
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("numgoroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("numgoroutine:", runtime.NumGoroutine())
	fmt.Println("final value: ", counter)
}
