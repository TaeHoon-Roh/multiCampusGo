package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dealer struct {
	card_num      int
	each_card_num int
	player_num    int
	card          []int // 매 게임 바뀜
	player        []Player
}

type Player struct {
	alive       bool
	money       int
	cardsumGame int
	win_count   int
	dropRound   int
}

func main() {
	dealer_ins := Dealer{20, 4, 2, nil, nil}
	// 처음 시작할때 참여자들 입장
	//player_ins := make([]Player,dealer_ins.player_num+1)
	// player 승률 및 자본금 초기화
	for inx := 1; inx < dealer_ins.player_num+1; inx++ {
		fmt.Println(inx)
		dealer_ins.player[inx].clear_player()
	}

	dealer_ins.make_CardSet(dealer_ins.card_num)
	// 100번의 게임
	for i := 0; i < 100; i++ {

		dealer_ins.shuffle_Card()
		//dealer_ins.give_card_to_player(dealer_ins.player_num,dealer_ins.each_card_num,dealer_ins.card)

		dealer_ins.give_card_to_player(dealer_ins.each_card_num, dealer_ins.card) // dealer.player[n].cardSum에 저장
		fmt.Println(dealer_ins.player[1].win_count)

		winner_once := dealer_ins.select_winner()

		for j := 1; j < dealer_ins.player_num+1; j++ {
			if j == winner_once {
				//player_ins[j].win_player()
				dealer_ins.player[j].money++
				dealer_ins.player[j].win_count++
			} else if j != 0 {
				//money_left := player_ins[j].lose_player()
				dealer_ins.player[j].money--
				if dealer_ins.player[j].money == 0 {
					dealer_ins.player[j].alive = false
					dealer_ins.player[j].dropRound = i
					dealer_ins.player_num--
				}
			}
		}
	}

	// 결과 출력
	for z := 1; z < len(dealer_ins.player); z++ {
		fmt.Printf("Player%d's Win Rate : %d %% / when dead : %d 번째 게임\n", z, dealer_ins.player[z].win_count, dealer_ins.player[z].dropRound)
	}
	fmt.Printf("(Draw Rate : %d %%)", dealer_ins.player[0].win_count)
}

func (p *Player) lose_player() int {
	p.money--
	return p.money
}

func (p *Player) win_player() {
	p.money++
	p.win_count++
}

func (p *Player) clear_player() {
	p.alive = true
	p.money = 10
	p.cardsumGame = 0
	p.win_count = 0
}

func (d *Dealer) select_winner() int {
	winner := 0
	max := 0
	each_player_result := make([]int, len(d.player))

	for i := range d.player {
		each_player_result[i] = d.player[i].cardsumGame % 10
		if max < each_player_result[i] {
			max, winner = each_player_result[i], i+1
		} else if max == each_player_result[i] {
			return 0
		}
	}
	return winner
}

func (d *Dealer) give_card_to_player(each_card_num int, card []int) {
	// player_card := make([]int,len(player)) // 한 게임에서 할당된 카드 저장하는 slice
	//card := makeCardSet()
	//check_exist := [len(card)]bool{false,}

	cardindex := 0
	//playerSum := make([]int,d.player_num+1)
	//save_playernum := []int{}

	for i := 0; i < each_card_num; i++ {
		for j := 1; j < d.player_num+1; j++ {
			//save_playernum[j] = in_player[j].number
			if d.player[j].alive == false {
				continue
			} else {
				d.player[j].cardsumGame += d.card[cardindex]
				fmt.Println(d.player[j].cardsumGame)
				//playerSum[j] += card[cardindex]
				cardindex++
			}
		}
	}
}

func (d *Dealer) shuffle_Card() {

	shuffledCard := make([]int, len(d.card))

	for i := range shuffledCard {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomN := r1.Intn(len(d.card))
		shuffledCard[i] = d.card[randomN]
		d.card = append(d.card[:randomN], d.card[randomN+1:]...)
	}
	d.card = shuffledCard
}

func (d *Dealer) make_CardSet(card_num int) {

	d.card = make([]int, card_num)
	for i := 0; i < 20; i++ {
		d.card[i] = (i + 1) % 10
		if d.card[i] == 0 {
			d.card[i] = 10
		}
	}

}
