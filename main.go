package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func heavyWork() int {
	time.Sleep(time.Second * 1)

	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	go func() {

		for i := 0; i < 100; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				c <- heavyWork()
			}()
		}

		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("program end successfully!!")
}
