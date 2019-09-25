package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("End Main")
	fmt.Println("hi")

	var wait sync.WaitGroup

	c1 := make(chan int)
	c2 := make(chan int)

	wait.Add(1)
	go func() {
		defer wait.Done()
		for {
			time.Sleep(time.Second)
			c1 <- 0
		}
	}()
	wait.Add(1)
	go func() {
		defer wait.Done()
		for {
			time.Sleep(time.Second * 2)
			c2 <- 1
		}
	}()
	wait.Add(1)
	go func() {
		defer wait.Done()
		for {
			select {
			case <-c1:
				fmt.Println("chan c1")
			case <-c2:
				fmt.Println("chan c2")

			}
		}
	}()

	wait.Wait()
}
