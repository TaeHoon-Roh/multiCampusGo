package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var winRate [3]int
	winRate[0] = 0
	winRate[1] = 0
	winRate[2] = 0
	for i := 0; i < 100; i++ {
		card := makeCard()
		mycard := shuffleCard(card)
		fmt.Println(mycard)

		var player_1, player_2 [2]int
		player_1[0] = mycard[0]
		player_2[0] = mycard[1]
		player_1[1] = mycard[2]
		player_2[1] = mycard[3]

		player_1_result, player_2_result, winner := compatition(player_1, player_2)
		fmt.Println(player_1_result, player_2_result, winner)
		winRate[winner]++
	}

	fmt.Println(winRate)

	fmt.Printf("Player1 Win Rate : %0.2f %%\n", float32(winRate[1]))
	fmt.Printf("Player2 Win Rate : %0.2f %%\n", float32(winRate[2]))
	fmt.Printf("Draw Rate : %0.2f %%\n", float32(winRate[0]))

}

func compatition(player_1 [2]int, player_2 [2]int) (int, int, int) {

	player_1_result := (player_1[0] + player_1[1]) % 10
	player_2_result := (player_2[0] + player_2[1]) % 10
	winner := 0
	if player_1_result > player_2_result {
		winner = 1
	} else if player_1_result == player_2_result {
		winner = 0
	} else {
		winner = 2
	}
	return player_1_result, player_2_result, winner
}

func shuffleCard(card [20]int) [20]int {

	var myCard [20]int
	fmt.Println(card)
	var index = 0
	for {
		flag := checkCard(card)
		if flag == 0 || index == 20 {
			break
		}

		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(20)
		if card[randomNumber] != 0 {
			myCard[index] = card[randomNumber]
			card[randomNumber] = 0
			//fmt.Println(myCard, card)
			index++
		} else {
			continue
		}

	}
	return myCard

}

func checkCard(card [20]int) int {

	flag := 0
	for _, temp := range card {
		flag = flag + temp
	}
	return flag
}

func makeCard() [20]int {

	var result [20]int
	for i := 0; i < 20; i++ {
		result[i] = (i + 1) % 10
		if result[i] == 0 {
			result[i] = 10
		}
	}
	return result
}
