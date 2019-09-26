package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	myListen, err := net.Listen("tcp", ":65000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		connect, err := myListen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ConnectHandler(connect)
	}

	defer func() {
		myListen.Close()
	}()
}

func ConnectHandler(connect net.Conn) {
	recvBuf := make([]byte, 4096) // receive buffer: 4kB
	userCount := 0
	for {
		n, err := connect.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				fmt.Println("connection is closed from client: ", connect.RemoteAddr().String())
				return
			}
			fmt.Println(err)
			return
		}
		data := string(recvBuf[:n])
		if 0 < n {
			fmt.Println(data)
		}

		if data == "add_user" {
			userCount++
		} else if data == "start" {
			if userCount > 5 {
				userCount = 5
			}
			startCardGameWithDealer(userCount)
		}
	}
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

func startCardGameWithDealer(userCount int) {
	// Create Dealer
	cardDealer := new(CardDealer)

	// Player Count.
	playerCount := userCount

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
