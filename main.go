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

	for _, byte := range dat {
		if 'A' <= byte && byte <= 'Z' || 'a' <= byte && byte <= 'z' || byte == ' ' {
			repDat = append(repDat, byte)
		}
	}

	//space로 단어별 절단
	oriStrArr := strings.Split(string(repDat), " ")

	resultArr := make([]map[string]int, 0)

	const ThreadCount int = 4

	splitCount := len(oriStrArr) / ThreadCount

	wait := sync.WaitGroup{}

	for len(oriStrArr) > 0 {
		var splitStrArr []string

		if len(oriStrArr) > splitCount {
			splitStrArr = oriStrArr[:splitCount]
			oriStrArr = oriStrArr[splitCount:]
		} else {
			splitStrArr = oriStrArr
			oriStrArr = make([]string, 0)
		}

		//Thread 생성 및 실행
		wait.Add(1)
		go func(splitStr []string) {
			defer wait.Done()

			countMap := make(map[string]int)

			for _, str := range splitStr {
				if str != "" {
					countMap[str]++
				}
			}

			resultArr = append(resultArr, countMap)
		}(splitStrArr)

	}

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

type WordObj struct {
	word  string
	count int
}
