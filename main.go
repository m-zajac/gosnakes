package main

import (
	"fmt"
	"log"
	"time"

	"github.com/m-zajac/gosnakes/snakes"
)

func main() {
	player := snakes.NewRandomPlayer("p1")
	board := snakes.Board{
		Size: [2]int{10, 10},
		Snakes: []snakes.Snake{
			snakes.Snake{
				ID:   player.ID(),
				Head: [2]int{5, 2},
				Tail: [][2]int{},
			},
		},
		Apples: [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}},
	}
	game, err := snakes.NewGame(board, player.ID(), player)
	if err != nil {
		log.Fatalf("got NewGame err: %v", err)
	}

	fmt.Println(game.Board().String())
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Millisecond)
		game.Next()
		fmt.Println(game.Board().String())
	}

}
