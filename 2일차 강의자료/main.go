package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dealer struct {
	card    []int
	name    string
	players []Player
	round   int
	draw    int
}

func (d *Dealer) makeCard() {
	d.card = make([]int, 20)
	for i := 0; i < 20; i++ {
		d.card[i] = (i + 1) % 10
		if d.card[i] == 0 {
			d.card[i] = 10
		}
	}
}

func (d *Dealer) cardShuffle() {

	myCard := make([]int, 20)

	for index := range myCard {
		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(d.card))
		myCard[index] = d.card[randomNumber]
		d.card = append(d.card[:randomNumber], d.card[randomNumber+1:]...)
	}

	d.card = myCard
}

func (d *Dealer) addPlayer(p Player) {
	d.players = append(d.players, p)
}

func (d *Dealer) startGame() {
	for index := range d.players {
		d.players[index].mycard = append(d.players[index].mycard, d.card[(index*2):(index*2)+2])
	}
	d.checkGame()

}

func (d *Dealer) checkGame() {
	playerResult := make([]int, 0, 5)
	for index := range d.players {
		playerResult = append(playerResult, (int(d.players[index].mycard[d.round][0]+d.players[index].mycard[d.round][1]))%10)
	}
	max := -1
	draw_check := 0
	max_index := 0
	for i := range playerResult {
		temp := playerResult[i]
		if temp > max {
			max = temp
			max_index = i
		} else if max == temp {
			draw_check++
		} else {
			continue
		}
	}

	if draw_check == 0 {
		d.players[max_index].winHit++
	} else {
		d.draw++
	}
	d.round++
	// fmt.Println(playerResult)
}

func (d *Dealer) printPlayerStatus() {
	for i := range d.players {
		fmt.Println("Player ", i, " WinHit : ", d.players[i].winHit)
	}
	fmt.Println("Round : ", d.round)
	fmt.Println("Draw Round: ", d.draw)
}

type Player struct {
	name      string
	age       int
	mycard    [][]int
	winHit    float32
	dropRound int
}

func (p *Player) receiveCard(receive_card []int) {
	p.mycard = append(p.mycard, receive_card)
}

func main() {

	dealer := Dealer{name: "dealer"}
	dealer.makeCard()

	player_1 := Player{name: "taehoon", age: 30}
	dealer.addPlayer(player_1)
	player_2 := Player{name: "roh", age: 30}
	dealer.addPlayer(player_2)
	fmt.Println(dealer)

	for i := 0; i < 100; i++ {
		dealer.cardShuffle()
		dealer.startGame()
	}
	dealer.printPlayerStatus()

}
