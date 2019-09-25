package main

import (
	"fmt"
	"sync"
	"time"
)

func myTest_1(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_1", buff)
}

func myTest_2(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_2", buff)
}

func myTest_3(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_3", buff)
}

func myTestIn(wg *sync.WaitGroup, mychannel chan string) {
	defer wg.Done()
	buff := "myTestIn"
	mychannel <- buff
}

func main() {
	var wait sync.WaitGroup
	defer func() {
		fmt.Println("main gorutine End!!")
	}()
	myChannel_1 := make(chan int)
	myChannel_2 := make(chan string)

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Func1")
		time.Sleep(time.Second)
		myChannel_1 <- 1
		close(myChannel_1)
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		val, flag := <-myChannel_2
		if !flag {
			fmt.Println("channel Close")
		} else {
			fmt.Println("func :", val, flag)
		}
	}()

	val, flag := <-myChannel_1
	if !flag {
		fmt.Println("channel Close")
	} else {
		fmt.Println("main :", val, flag)
	}
	myChannel_2 <- "HIHI"
	close(myChannel_2)

	wait.Wait()
}
