package main

import (
	"fmt"
	"strconv"
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

func myTestIn(wg *sync.WaitGroup, mychannel chan<- string) {
	defer wg.Done()
	buff := "myTestIn"
	mychannel <- buff
}

func main() {
	defer fmt.Println("main gorutine End!!")
	var s1, s2 string
	var wait sync.WaitGroup
	myChannel := make(chan string, 10)
	fmt.Println(myChannel)
	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Input Data!")
		// n, b := fmt.Scan(&s1, &s2)
		// fmt.Println(n, b)
		fmt.Println(s1, s2)
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			myChannel <- ("hihi " + strconv.Itoa(i))

		}
		close(myChannel)
		// myChannel <- "hihi"
	}()

	wait.Add(1)
	go myTest_1(&wait, myChannel)
	wait.Add(1)
	go myTest_2(&wait, myChannel)
	wait.Add(1)
	go myTestIn(&wait, myChannel)
	wait.Add(1)
	go myTestIn(&wait, myChannel)
	wait.Add(1)
	go myTestIn(&wait, myChannel)
	wait.Add(1)
	go myTest_3(&wait, myChannel)

	for i := range myChannel {
		println(i)
	}
	wait.Wait()

}
