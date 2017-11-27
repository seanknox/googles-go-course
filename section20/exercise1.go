package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("CPU(s): ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		fmt.Println("This is one goroutine")
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		fmt.Println("This is another goroutine")
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()
}
