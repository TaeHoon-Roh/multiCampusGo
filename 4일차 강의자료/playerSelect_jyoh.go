package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Dealer struct {
	name    string
	card    []int
	players []player
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

func main() {
	defer fmt.Println("End Main")
	fmt.Println("hi")

	dealer := Dealer{name: "dealer"}
	dealer.makeCard()

	player_1 := Player{name: "taehoon", age: 30}
	dealer.addPlayer(player_1)
	player_2 := Player{name: "roh", age: 30}
	dealer.addPlayer(player_2)
	fmt.Println(dealer)

	var wait sync.WaitGroup

	//c1 := make(chan int)
	//c2 := make(chan int)

	c_in := make(chan int, 5)
	c_out := make(chan int, 5)

	for i := 0; i < len(dealer.players); i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for {
				c_in <- i
				temp := <-c_out
				time.Sleep(time.Second)
				fmt.Println("routine ", i, temp)
			}
		}()
	}

	wait.Add(1)

	go func() {
		defer wait.Done()
		for {

			select {
			case <-time.Tick(time.Second * 3):
				fmt.Println("timeTick")

			}

			fmt.Println("Delear : select end!")

			temp := len(c_in)
			fmt.Println("channel result : ", temp)
			for i := 0; i < temp; i++ {
				mytemp := <-c_in
				c_out <- mytemp
			}
		}
	}()

	wait.Wait()
}
