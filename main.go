package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	// startCardGame()

	// card := makeCard()
	// fmt.Println("origin card : ", card)
	// result := shuffleCard(card)
	// fmt.Println("suffled card : ", result)

	// a := make(map[int]int)
	// a[1] = 100
	// val, flag := a[1]
	// fmt.Println("val : ", val, ", flag : ", flag)
	// fmt.Println("a : ", a[1])

	// startGameWithDealer()

	// 파일을 읽어옴.
	dat, err := ioutil.ReadFile("C:\\workspace_go\\text\\test02.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// 스레드가 처리해야 할 남은 길이.
	remainWordLength := len(dat)

	// 스레드 wait 그룹.
	var wait sync.WaitGroup

	// 결과 값을 저장할 맵.
	storyMap := make(map[string]int)

	// 하나의 스레드가 처리해야할 데이터 사이즈.
	processSize := 1000

	// 스레드 생성 및 실행.
	i := 0
	for {
		wait.Add(1)
		datPart := make([]byte, 0)
		startIndex := i * processSize
		endIndex := startIndex + processSize
		if remainWordLength < processSize {
			endIndex = startIndex + remainWordLength
		}
		datPart = append(datPart, dat[startIndex:endIndex]...)
		remainWordLength -= processSize
		i++
		go getCountOfWord(datPart, storyMap, &wait)
		if remainWordLength <= 0 {
			break
		}
	}

	wait.Wait()
	fmt.Println("End of Program!!!")
	for i, v := range storyMap {
		fmt.Println(i, " : ", v)
	}

	fmt.Println("============================")
	fmt.Println("total length : ", len(dat))
	fmt.Println("Process Size : ", processSize)
	fmt.Println("total thread count : ", i)

}

func getCountOfWord(dat []byte, storyMap map[string]int, wait *sync.WaitGroup) {
	defer wait.Done()

	fmt.Println("Start getCountOfWord : ", len(dat))

	// 공백, A~Z, a~z
	for i := 0; i < len(dat); i++ {
		v := dat[i]
		if (v == ' ' || ('A' <= v && v <= 'Z') || ('a' <= v && v <= 'z')) == false {
			// fmt.Println("remove : ", v, " : ", string(v))
			dat = append(dat[:i], dat[i+1:]...)
			i--
		}
	}

	// 파일 내용을 저장할 문자열 변수.
	story := string(dat)
	// fmt.Println(story)

	var splitStory []string
	splitStory = strings.Split(story, " ")

	for _, v := range splitStory {
		if v == " " || v == "" {
			continue
		}
		count := storyMap[v]
		if count == 0 {
			storyMap[v] = 1
		} else {
			storyMap[v]++
		}
	}
}

func mergeCountOfWord(maps ...map[string]int) {

}

type CardDealer struct {
	player   []Player
	cardList []int
}

type Player struct {
	coin            int
	cardResult      int
	winGameCount    int
	bankruptcyRound int
}

func startGameWithDealer() {
	// Create Dealer
	cardDealer := new(CardDealer)

	// Player Count.
	playerCount := 5

	defaultCoin := 10

	// Assign player to dealer
	cardDealer.player = make([]Player, playerCount)
	for i := range cardDealer.player {
		cardDealer.player[i].coin = defaultCoin
	}

	drawGameRoundList := make([]int, 0)

	totalGameCount := 100
	drawGameCount := 0

	for gameRound := 0; gameRound < totalGameCount; gameRound++ {
		fmt.Println("===== Start : ", gameRound+1)

		// Ready Card
		card := makeCard()
		fmt.Println("origin card : ", card)

		cardDealer.cardList = shuffleCard(card)
		fmt.Println("suffled card : ", cardDealer.cardList)

		// init player gamaResult.
		for i := range cardDealer.player {
			cardDealer.player[i].cardResult = -1
		}

		// Player Count
		for i := range cardDealer.player {
			// Divide And Get Result.
			if cardDealer.player[i].coin > 0 {
				cardDealer.player[i].cardResult = (cardDealer.cardList[i*2] + cardDealer.cardList[(i*2)+1]) % 10
			}
		}

		winValue := 0
		winIndex := 0
		for i := range cardDealer.player {
			if cardDealer.player[i].cardResult > winValue {
				winValue = cardDealer.player[i].cardResult
				winIndex = i
			}
		}

		checkDrawGame := 0
		for i := range cardDealer.player {
			if cardDealer.player[i].cardResult == winValue {
				checkDrawGame++
			}
		}

		if checkDrawGame > 1 {
			drawGameCount++
			drawGameRoundList = append(drawGameRoundList, gameRound)
			fmt.Println("Draw Game : ", gameRound)
		} else {
			for i := range cardDealer.player {
				if cardDealer.player[i].coin > 0 {
					if winIndex != i {
						cardDealer.player[i].coin--
						cardDealer.player[winIndex].coin++

						if cardDealer.player[i].coin <= 0 {
							cardDealer.player[i].bankruptcyRound = gameRound
						}
					} else {
						cardDealer.player[winIndex].winGameCount++
					}
				}
			}
			fmt.Println("Result : ", cardDealer.player)
		}
	}

	fmt.Println("========================")
	fmt.Println("Total Game : ", totalGameCount, ", defaultCoin : ", defaultCoin)
	for i := range cardDealer.player {
		fmt.Println("Player ", i, " : coin (", cardDealer.player[i].coin, "), win(", cardDealer.player[i].winGameCount, "), bankruptcyRound(", cardDealer.player[i].bankruptcyRound, ")")
	}
	fmt.Println("Draw Game : ", drawGameCount, " : ", drawGameRoundList)
}

func increase(number *int) {
	*number++
}

func startCardGame() {

	// 카드 값 초기화.
	cardList := make([]int, 20)
	for i := 0; i < 20; i++ {
		cardValue := (i + 1)
		if cardValue > 10 {
			cardValue -= 10
		}
		cardList[i] = cardValue
	}
	fmt.Println("카드 정보 : ", cardList)

	totalGameCount := 100
	userCount := 5 // 5 를 넘어갈 수 없음
	selectedCardSize := userCount * 2

	userWin := make([]int, userCount)
	noWinner := 0

	for gameCount := 0; gameCount < totalGameCount; gameCount++ {
		fmt.Println("===== Start ", gameCount+1, "th")

		// 매번 새로운 랜덤 값 생성을 위해 사용.
		rand.Seed(time.Now().UnixNano())

		// 카드를 선택.
		selectedCardList := make([]int, selectedCardSize)
		for i := 0; i < selectedCardSize; i++ {
			duplicated := 0
			pickedCard := cardList[pickCard(20)]
			for _, value := range selectedCardList {
				if value == pickedCard {
					duplicated++
				}
				if duplicated >= 2 {
					// fmt.Println("2장이상 겹치는 경우 발견!!", value, " : ", pickedCard)
					i--
					break
				} else {
					selectedCardList[i] = pickedCard
				}
			}
		}
		fmt.Println("뽑은 카드 : ", selectedCardList)

		// 카드 계산.
		gameResultList := make([]int, userCount)
		for i := 0; i < userCount; i++ {
			gameResultList[i] = (selectedCardList[i*2] + selectedCardList[(i*2)+1]) % 10
		}

		winValue := 0
		winIndex := 0
		for index, value := range gameResultList {
			if value >= winValue {
				winValue = value
				winIndex = index
			}
		}

		checkDrawGame := 0
		for i := 0; i < userCount; i++ {
			if gameResultList[i] == winValue {
				checkDrawGame++
			}
		}

		fmt.Println("게임 결과 : ", gameResultList)
		if checkDrawGame > 1 {
			noWinner++
			fmt.Println("Draw Game!!!")
		} else {
			userWin[winIndex] = userWin[winIndex] + 1
			fmt.Println("Current Result : ", userWin)
		}
	}

	fmt.Println("=================================")
	fmt.Println("Total Game : ", totalGameCount)
	fmt.Println("Game Result: ", userWin)
	fmt.Println("Draw Game Count : ", noWinner)
}

func pickCard(cardSize int) int {
	return rand.Intn(cardSize)
}

func shuffleCard(card []int) []int {

	myCard := make([]int, len(card))
	for index := range myCard {
		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(card))
		myCard[index] = card[randomNumber]
		card = append(card[:randomNumber], card[randomNumber+1:]...)
	}
	return myCard
}

func makeCard() []int {
	result := make([]int, 20)
	for i := 0; i < 20; i++ {
		result[i] = (i + 1) % 10
		if result[i] == 0 {
			result[i] = 10
		}
	}
	return result
}

func checkCard(card []int) int {
	flag := 0
	for _, temp := range card {
		flag = flag + temp
	}
	return flag
}
