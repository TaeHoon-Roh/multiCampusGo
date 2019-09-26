//카드 게임 - 구조체 사용

package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	dealer := Dealer{}

	connChan := make(chan net.Conn)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go dealer.initPlayers(&wg, connChan)
	wg.Add(1)
	go dealer.serverOpen(&wg, connChan)

	playGame(&dealer)

	wg.Wait()
}

func playGame(dealer *Dealer) {
	const DelaySeconds = 1

	for {
		if len(dealer.players) >= 2 {
			dealer.passCard()
			dealer.compatition()

			for i, player := range dealer.players {
				fmt.Printf("Player%d Win Rate : %0.2f %%\n", i, player.getWinRate())
				fmt.Printf("Player%d win : %d, lose : %d, draw : %d, gameCount : %d, coin : %d\n", i, player.win, player.lose, player.draw, player.gameCount, player.coin)
			}

		} else {
			fmt.Printf("Wait for players(%d)\n", len(dealer.players))
		}

		time.Sleep(time.Second * DelaySeconds)
	}
}

//Player
type Player struct {
	name                       int
	win, lose, draw, gameCount int
	coin                       int
	card1, card2               int
	connect                    net.Conn
}

func (player Player) getWinRate() float32 {
	winRate := float32(player.win) / float32(player.gameCount) * 100
	return winRate
}

func (player Player) getScore() int {
	return (player.card1 + player.card2) % 10
}

//Dealer
type Dealer struct {
	cards   []int
	players []Player
}

func (dealer *Dealer) initPlayers(wg *sync.WaitGroup, conchan <-chan net.Conn) {
	defer wg.Done()

	const DefCoin int = 10
	dealer.players = make([]Player, 0)

	for {
		connect := <-conchan

		fmt.Println("make players ", len(dealer.players))

		//Player 생성
		dealer.players = append(dealer.players, Player{name: len(dealer.players) + 1, coin: DefCoin, connect: connect})

		fmt.Println("make player")
	}
}

func (dealer *Dealer) serverOpen(wg *sync.WaitGroup, conchan chan<- net.Conn) {
	myListen, err := net.Listen("tcp", ":5554")
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

		conchan <- connect
	}

	defer func() {
		myListen.Close()
		wg.Done()
	}()
}

func (dealer *Dealer) shuffleCards() {
	shuppleCards := make([]int, 20)
	cards := dealer.cards

	for i := range cards {
		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(cards))

		shuppleCards[i] = cards[randomNumber]

		cards = append(cards[:randomNumber], cards[randomNumber+1:]...)
	}

	dealer.cards = shuppleCards
}

func (dealer *Dealer) makeCard() {
	cards := make([]int, 20)
	for i := 0; i < 20; i++ {
		cards[i] = (i + 1) % 10
		if cards[i] == 0 {
			cards[i] = 10
		}
	}

	dealer.cards = cards
}

func (dealer *Dealer) passCard() {
	for i := 0; i < 100; i++ {
		dealer.makeCard()
		dealer.shuffleCards()

		for j := range dealer.players {
			dealer.players[j].card1 = dealer.cards[0]
			dealer.players[j].card2 = dealer.cards[1]

			dealer.cards = dealer.cards[2:]
		}
	}
}

func (dealer *Dealer) compatition() {
	maxScore := -1
	winner := -1
	livePlayerCount := 0
	livePlayerName := -1

	for i := range dealer.players {
		var score int = dealer.players[i].getScore()

		if dealer.players[i].coin == 0 {
			fmt.Printf("Player %d is dead\n", i)
			continue
		} else {
			livePlayerCount++
			livePlayerName = dealer.players[i].name
		}

		if score > maxScore {
			winner = i
			maxScore = score
		} else if score == maxScore {
			//draw
			winner = -1
			maxScore = score
		}
	}

	//생존 Player가 하나면 게임을 진행하지 않음
	if livePlayerCount <= 1 {
		fmt.Printf("Live player is Only one!! (%d)\n", livePlayerName)
		return
	}

	for i := range dealer.players {
		if dealer.players[i].coin == 0 {
			continue
		}

		dealer.players[i].gameCount++

		if winner == -1 {
			dealer.players[i].draw++
		} else if winner == i {
			dealer.players[i].win++
			dealer.players[i].coin = dealer.players[i].coin + livePlayerCount - 1
		} else if winner != i {
			dealer.players[i].lose++
			dealer.players[i].coin--
		}

	}

	fmt.Println("winner : ", winner)

}
