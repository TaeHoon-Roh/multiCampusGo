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

bookData, err := ioutil.ReadFile("/Users/user/go/src/WORKSPACE_GO/src/main/test_1.txt")
if err != nil {
log.Fatal(err)
}

words := filterRegex(string(bookData))

//========== 미완 =============
//인덱스를 특정 단위로 나눠서 멀티 쓰레드....
// 쪼갠 배열 만들기??

subWords := sliceWords(2,words)

var wait sync.WaitGroup
var subCounter []map[string]int

go func() {
for i:= range subWords {
wait.Add(1)
subCounter = append(subCounter,countWords(subWords[i]))
}
}()

counter := mergeMap(subCounter)
//===========================

// IF, ascii로 필터링 할 경우 아래 코드 사용
//bookString := removeSP(bookData)
//words := strings.Split(bookString, " ")

//counter := countWords(words) //[:len(words)/100])


for key, val := range counter {
fmt.Println(key, val)
}
wait.Wait()
}

func mergeMap(subCounter []map[string]int) map[string]int {
counter := map[string]int{}
for _, subValue := range subCounter {
for k, v := range subValue {
counter[k] = counter[k] + v
}
}
return counter
}


func sliceWords (threadNum int, words []string) [][]string {
splitNum := len(words)/threadNum
var subWords [][]string

for i:=0; i<threadNum; i++ {
subWords[i] = words[:splitNum]
words = words[splitNum:]
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