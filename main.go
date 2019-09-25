package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var mt = new(sync.Mutex)

func main() {

	defer fmt.Println("Finish=================")
	fmt.Println("Start=================")

	var ws sync.WaitGroup
	var tokens []string

	// read text file
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	chan1 := make(chan int)
	chan2 := make(chan int)

	// create map
	tmap := make(map[string]int)

	ws.Add(1)
	go func() {
		defer ws.Done()
		fmt.Println(">> collect alphabet characters from text <<")
		sz := len(data)
		// convert special character to blank
		for i := 0; i < sz; i++ {
			if !(data[i] >= 'A' && data[i] <= 'Z') && !(data[i] >= 'a' && data[i] <= 'z') && data[i] != ' ' {
				data[i] = ' '
			}
		}
		//fmt.Println(string(data))
		chan1 <- 0
	}()

	ws.Add(1)
	go func() {
		defer ws.Done()
		n := <-chan1
		fmt.Println("from chan1", n)
		fmt.Println(">> parse words & tokenize them from text <<")
		tokens = parse(strings.ToLower(string(data)))
		chan2 <- 1
	}()

	ws.Add(1)

	go func() {
		defer ws.Done()
		n := <-chan2
		fmt.Println("from chan2", n)
		fmt.Println(">> count word-token <<")
		// run go-routine
		loopcnt := 2
		for i := 0; i < loopcnt; i++ {
			ws.Add(1)
			// count word in text
			go count(&ws, tokens, tmap, i, loopcnt)
		}
	}()

	ws.Wait()

	// print out map-data
	tot := 0
	wc := 0
	for key, value := range tmap {
		fmt.Println(key, value)
		tot += value
		wc++
	}
	fmt.Println("number of word-token", len(tokens))
	fmt.Println("total word count in map", tot)
	fmt.Println("number of word", wc)

}

func parse(data string) []string {
	return strings.Split(data, " ")
}

func count(ws *sync.WaitGroup, tokens []string, tmap map[string]int, group int, numGo int) {

	procCnt := 0
	defer func() {
		fmt.Printf("Go Routine Finish(%d) - procCnt = %d\n", group, procCnt)
		ws.Done()
	}()

	fmt.Printf("Go Routine Start(%d)\n", group)

	for idx, value := range tokens {
		if (idx % numGo) == group {
			procCnt++
			mt.Lock()
			if _, ok := tmap[value]; ok {
				tmap[value]++
			} else {
				tmap[value] = 1
			}
			mt.Unlock()
		}
	}
}
