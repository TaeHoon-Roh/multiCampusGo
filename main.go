package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var detMap = make(map[string]int)
	// var detMap_1 = make(map[string]int)
	// var detMap_2 = make(map[string]int)
	var detByte = make([]byte, 0)
	// var wait sync.WaitGroup
	det, err := ioutil.ReadFile("C:\\workspace_go\\src\\test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(len(det))
	for _, val := range det {
		// fmt.Println(val)
		// if 'A' <= val && val <= 'Z' || 'a' <= val <= 'z' || val == ' ' {
		if val >= 'A' && val <= 'Z' || val >= 'a' && val <= 'z' || val == ' ' {
			detByte = append(detByte, val)
		}
	}
	// fmt.Println(detByte)
	detSplit := strings.Split(string(detByte), " ")
	// fmt.Println(detSplit)
	det_1 := detSplit[:len(detSplit)/2]
	fmt.Println(det_1)
	det_2 := detSplit[len(detSplit)/2:]
	fmt.Println(det_2)

	// wait.Add(1)

	myfunc := func(str []string) (temp map[string]int) {
		// defer wait.Done()
		for _, val := range str {
			key := val
			_, isKey := temp[key]
			// fmt.Println(val, isKey)
			if isKey != true {
				temp[key] = 1
			} else {
				temp[key]++
			}
		}
		return temp
	}

	detMap_1 := myfunc(det_1)
	detMap_2 := myfunc(det_2)

	fmt.Println("Main Tread!")
	// wait.Wait()

	// for key, _ := range detMap {
	fmt.Println(detMap_1)
	fmt.Println(detMap_2)
	fmt.Println(detMap)
	// }
	// flag := strings.Contains(detSpli, "")
}
