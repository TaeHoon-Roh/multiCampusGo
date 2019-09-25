package main

import (
	"fmt"
	"strconv"
	"sync"
)

func myThreadTest(wait *sync.WaitGroup, number int, temp string, buffer ...int) {
	defer wait.Done()
	fmt.Println(buffer)
	buffer[0] = number * 1
	buffer[1] = number * 2
	fmt.Println(temp+strconv.Itoa(number)+" out Thread", buffer)
}

func change(buffer []int) {
	buffer[0] = 10
	buffer[1] = 20
}

func main() {
	temp := "Hello "
	buffer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	defer func() {
		fmt.Println(buffer)
		fmt.Println("Main End!")
	}()
	var wait sync.WaitGroup

	change(buffer)
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go myThreadTest(&wait, i, temp, buffer...)
	}

	fmt.Println("Main Thread!")
	wait.Wait()

	// var wait sync.WaitGroup
	// myChanner := make(chan int)

	// wait.Add(1)
	// go func() {
	// 	defer wait.Done()
	// 	fmt.Println("Function Wait...")
	// 	temp := <-myChanner
	// 	fmt.Println("Channer val is ", temp)

	// }()

	// time.Sleep(time.Second)
	// myChanner <- 1
	// close(myChanner)
	// wait.Wait()
}
