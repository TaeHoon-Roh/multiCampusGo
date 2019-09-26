package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sync"
)

func main() {

	bookData, err := ioutil.ReadFile("C:\\Users\\student\\Desktop\\test_1.txt")

	if err != nil {
		log.Fatal(err)
	}

	// regex 기반 data filtering
	words := filterRegex(string(bookData))

	// 전체 데이터를 특정 단위(Thread 2개 예제)로 slicing
	subWords := sliceWords(2, words)

	// slicing 된 각각의 map 결과 저장하는 slice
	var subCounter []map[string]int

	var wait sync.WaitGroup

	// subslice 별 count 계산 (각각 결과를 map으로 받아옴)
	for i := range subWords {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			subCounter = append(subCounter, countWords(subWords[i]))
		}(i)
	}

	//===========================
	// IF, ascii로 필터링 할 경우 아래 코드 사용 (본 코드는 regex 사용)
	//bookString := removeSP(bookData)
	//words := strings.Split(bookString, " ")
	//===========================

	// 각각의 결과를 저장했던 map을 merge
	counter := mergeMap(subCounter)
	for key, val := range counter {
		fmt.Println(key, val)
	}
	wait.Wait()
}

// 슬라이스 별 생성된 map을 merge
func mergeMap(subCounter []map[string]int) map[string]int {
	counter := map[string]int{}
	for _, subValue := range subCounter {
		for k, v := range subValue {
			counter[k] = counter[k] + v
		}
	}
	return counter
}

func sliceWords(threadNum int, words []string) [][]string {
	splitSize := len(words) / threadNum
	var subWords [][]string

	for i := 0; i < threadNum; i++ {
		//for len(words) > splitSize {
		subWords[i] = words[:splitSize]
		//subWords = append(subWords,words[:splitSize])
		words = words[splitSize:]
	}
	return subWords
}

//func countWords(wait *sync.WaitGroup,words []string) map[string]int {
func countWords(words []string) map[string]int {
	//defer wait.Done()

	counter := map[string]int{}
	for _, word := range words {
		counter[word]++
	}
	return counter
}

// ASCII filtering
func removeSP(org_bookString []byte) string {

	bookString := ""
	for i := range org_bookString {
		if org_bookString[i] == 32 || (64 < org_bookString[i] && org_bookString[i] < 91) || (96 < org_bookString[i] && org_bookString[i] < 123) {
			bookString += string(org_bookString[i])
		}
	}

	return bookString
}

// REGEX filtering
func filterRegex(bookData string) []string {
	re, _ := regexp.Compile(`[a-zA-Z]*[a-zA-Z]`)
	return re.FindAllString(string(bookData), -1)
}
