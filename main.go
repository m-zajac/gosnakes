package main

import (
	"fmt"
	"log"
	"time"

	"github.com/m-zajac/gosnakes/snakes"
	"github.com/m-zajac/gosnakes/storage"
)

func main() {
	// db test
	db, err := storage.NewDB("./data")
	if err != nil {
		log.Fatalf("new db err: %v", err)
	}
	db.Write("test", "test", "entry2")
	var dbRes string
	db.Read("test", "test", &dbRes)
	fmt.Printf("db res: %s\n", dbRes)

	// game test
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
