package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Max Procs:", runtime.GOMAXPROCS(8))

	const maxRuns = 100
	var wg sync.WaitGroup
	wg.Add(maxRuns)

	for i := 0; i < maxRuns; i++ {
		go func() {
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("End value: ", counter)

}
