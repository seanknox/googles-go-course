package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	const maxRuns = 100
	var wg sync.WaitGroup
	wg.Add(maxRuns)

	var mu sync.Mutex

	for i := 0; i < maxRuns; i++ {
		go func() {
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			mu.Lock()
			v := counter
			// runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End value: ", counter)

}
