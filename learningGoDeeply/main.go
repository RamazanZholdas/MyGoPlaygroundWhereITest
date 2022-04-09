package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go work(&wg)
	wg.Wait()
	fmt.Println("main end")
	fmt.Println("Time passed", time.Since(now))
}

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work")
	time.Sleep(time.Second)
	fmt.Println("work end")
}
