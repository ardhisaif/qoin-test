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
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)

	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)
	diceGame(numPlayers, numDice)
}

type player struct {
	number int
	point  int
	switchDice int
	dice   []int
}

func diceGame(numPlayer int, numDice int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	players := make([]*player, numPlayer)

	fmt.Println("===========================")
	fmt.Println("Giliran 1 lempar dadu")
	for i := 0; i < numPlayer; i++ {
		players[i] = &player{
			number: i + 1,
			point:  0,
			dice:   make([]int,0),
		}

		fmt.Printf("pemain #%d (%d) :", players[i].number, players[i].point)


		for j := 0; j < numDice; j++ {
			num := r1.Intn(6) + 1
			fmt.Print(num, ",")

			if num == 1 {
				players[i].switchDice += 1
			}else if num == 6 {
				players[i].point++
			}else {
				players[i].dice = append(players[i].dice, num)
			}
		}

		fmt.Println()
	}

	for i := 0; i < numPlayer; i++ {

		var nextPlayer int
		if i == 0 {
			nextPlayer = 1
		} else if i == numPlayer-1 {
			nextPlayer = 0
		} else {
			nextPlayer = i + 1
		}

		for j := players[i].switchDice; j > 0 ; j-- {
			players[nextPlayer].dice = append(players[nextPlayer].dice, 1)
		}
	}

	fmt.Println("Setelah evaluasi:")
	for i := 0; i < numPlayer; i++ {
		
		fmt.Printf("pemain #%d (%d) :", players[i].number, players[i].point)
		fmt.Println(arrToString(players[i].dice))
	}
}

func arrToString(arr []int) string {
	var result string
	for i := range arr {
		text := strconv.Itoa(arr[i])
		result += text + ","
	}
	return result
}