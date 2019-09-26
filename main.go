package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

func main() {
	run()
}

func run() {
	dat, err := ioutil.ReadFile("C:\\workspace_go\\src\\main\\test_2.txt")

	if err != nil {
		fmt.Println("Error!!")
		fmt.Println(err)
	}

	//사용할 문자만 저장
	repDat := make([]byte, 0)

	wait := sync.WaitGroup{}

	for _, byte := range dat {
		if 'A' <= byte && byte <= 'Z' || 'a' <= byte && byte <= 'z' || byte == ' ' {
			repDat = append(repDat, byte)
		}
	}

	//space로 단어별 절단
	oriStrArr := strings.Split(string(repDat), " ")

	const ThreadCount int = 2

	splitCount := len(oriStrArr) / ThreadCount

	chanStr1 := make(chan string)
	chanStr2 := make(chan string)

	strings1 := oriStrArr[:splitCount]
	strings2 := oriStrArr[splitCount:]

	wait.Add(1)
	go strChanSend(&wait, chanStr1, &strings1)
	wait.Add(1)
	go strChanSend(&wait, chanStr2, &strings2)

	resultArr := make([]map[string]int, 0)

	wait.Add(1)
	go strChanReceive(&wait, chanStr1, &resultArr)
	wait.Add(1)
	go strChanReceive(&wait, chanStr2, &resultArr)

	//모든 Thread 실행 완료까지 대기
	wait.Wait()

	//Thread 결과 값을 하나의 Map 으로 병합
	totMap := make(map[string]int)
	totCnt := 0

	for _, countMap := range resultArr {
		for k, v := range countMap {
			totMap[k] = totMap[k] + v
			totCnt = totCnt + v
		}
	}

	//정렬을 위해 구조체 WordObj 배열을 이용
	wordObjArr := make([]WordObj, 0)

	for k, v := range totMap {
		wordObjArr = append(wordObjArr, WordObj{word: k, count: v})
	}

	//sort
	sort.Slice(wordObjArr, func(i, j int) bool {
		//오름차순
		return wordObjArr[i].count < wordObjArr[j].count
	})

	for _, wordObj := range wordObjArr {
		fmt.Println(wordObj.word, wordObj.count)
	}

	fmt.Printf("Total Count : %d", totCnt)
}

func strChanSend(wait *sync.WaitGroup, sendChan chan<- string, splitStrs *[]string) {
	defer wait.Done()

	for _, str := range *splitStrs {
		if str != "" {
			sendChan <- str
		}
	}

	close(sendChan)
}

func strChanReceive(wait *sync.WaitGroup, receiveChan <-chan string, resultArr *[]map[string]int) {
	defer wait.Done()

	fmt.Println("start", wait)

	countMap := make(map[string]int)

	for chanStr := range receiveChan {
		countMap[chanStr]++
	}

	fmt.Println("count thread end")

	*resultArr = append(*resultArr, countMap)
}

type WordObj struct {
	word  string
	count int
}
