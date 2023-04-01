package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	var numPlayers, numDice int

	// Meminta input dari user untuk jumlah pemain dan dadu
	// fmt.Print("Masukkan jumlah pemain: ")
	// fmt.Scan(&numPlayers)

	// fmt.Print("Masukkan jumlah dadu: ")
	// fmt.Scan(&numDice)
	numPlayers = 3
	numDice = 4
	players := make([]*player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &player{
			number: i + 1,
			isActive: true,
			dice:   make([]int,numDice),
		}
	}
	diceGame(players)
}

type player struct {
	number int
	point  int
	amountDice int
	switchDice int
	isActive bool
	dice   []int
}

var round int

func diceGame(players []*player) {
	for {

		round += 1
	
		fmt.Println("===========================")
		fmt.Printf("Giliran %d lempar dadu \n", round)
	
		startGame(players)
		
		switchI(players)
	
		fmt.Println("Setelah evaluasi:")
		playerAmount := eval(players)
		fmt.Println(playerAmount)
		if playerAmount < 2 {
			break
		}


		// for i := 0; i < numPlayer; i++ {
		// 	fmt.Println(players[i])
		// }
	}

}


func startGame(players []*player) {
	for i := 0; i < len(players); i++ {

		fmt.Printf("pemain #%d (%d) :", players[i].number, players[i].point)

		rollDice(len(players[i].dice), i, players)

		fmt.Println()
	}
}


func rollDice(amountDice int, i int, players []*player)  {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var temporaryDice []int
	for j := 0; j < amountDice; j++ {
		num := r1.Intn(6) + 1
		fmt.Print(num, ",")

		if num == 1 {
			players[i].switchDice += 1
		}else if num == 6 {
			players[i].point++
		}else {
			temporaryDice = append(temporaryDice, num)
		}
	}
	players[i].dice = temporaryDice
}

func switchI(players []*player)  {
	fmt.Println(players[0], players[1], players[2])
	for i := 0; i < len(players); i++ {
		var nextPlayer int
		if i == 0 {
			nextPlayer = 1
		} else if i == len(players)-1 {
			nextPlayer = 0
		} else {
			nextPlayer = i + 1
		}

		for {
			if len(players[nextPlayer].dice) == 0 {
				if nextPlayer == len(players)-1{
					nextPlayer = 0
				}else{
					nextPlayer ++ 
				}
			}else {
				break
			}
		}
		fmt.Println(nextPlayer)


		for j := players[i].switchDice; j > 0 ; j-- {
			players[nextPlayer].dice = append(players[nextPlayer].dice, 1)
		}
		players[i].switchDice = 0

	}
}


func eval(players []*player) int {
	var playerAmount int
	for i := 0; i < len(players); i++ {
		
		if len(players[i].dice) == 0{
			players[i].isActive = false
		}

		if len(players[i].dice) == 0{
			fmt.Printf("pemain #%d (%d) :", players[i].number, players[i].point)
		}else{
			fmt.Printf("pemain #%d (%d) :", players[i].number, players[i].point)
			fmt.Println(arrToString(players[i].dice))
		}
		if len(players[i].dice) > 0 {
			playerAmount ++
		}
		amount := &players[i].dice
		players[i].amountDice = len(*amount)
	}
	return playerAmount
}


func arrToString(arr []int) string {
	var result string
	for i := range arr {
		text := strconv.Itoa(arr[i])
		result += text + ","
	}
	return result
}