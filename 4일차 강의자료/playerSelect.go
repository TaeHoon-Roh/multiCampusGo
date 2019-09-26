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

	//c1 := make(chan int)
	//c2 := make(chan int)

	c_in := make(chan int, 5)
	c_out := make(chan int, 5)

	wait.Add(1)
	go func(myChannel chan int) {
		defer wait.Done()
		for {
			myChannel <- 1
			temp := <-c_out
			time.Sleep(time.Second)
			fmt.Println("routine 1 : ", temp)
		}
	}(c_in)

	wait.Add(1)
	go func(myChannel chan int) {
		defer wait.Done()
		for {
			c_in <- 2
			temp := <-c_out
			time.Sleep(time.Second)
			fmt.Println("routine 2 : ", temp)

		}
	}(c_in)
	wait.Add(1)

	go func() {
		defer wait.Done()
		for {

			select {
			case <-time.Tick(time.Second * 3):
				fmt.Println("timeTick")

			}

			fmt.Println("Delear : select end!")

			temp := len(c_in)
			fmt.Println("channel result : ", temp)
			for i := 0; i < temp; i++ {
				mytemp := <-c_in
				c_out <- mytemp
			}

		}
	}()

	wait.Wait()
}
