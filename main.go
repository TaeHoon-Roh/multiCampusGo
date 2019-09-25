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

	// read text file
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// get length of data
	sz := len(data)

	// convert special character to blank
	for i := 0; i < sz; i++ {
		if !(data[i] >= 'A' && data[i] <= 'Z') && !(data[i] >= 'a' && data[i] <= 'z') && data[i] != ' ' {
			data[i] = ' '
		}
	}
	fmt.Println(string(data))

	tokens := parse(strings.ToLower(string(data)))
	//tokens := parse(string(data))

	// create map
	tmap := make(map[string]int)

	// run go-routine
	loopcnt := 2
	for i := 0; i < loopcnt; i++ {
		ws.Add(1)
		// count word in text
		go count(&ws, tokens, tmap, i, loopcnt)
	}

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
