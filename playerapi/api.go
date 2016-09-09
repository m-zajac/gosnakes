package playerapi

import (
	"log"
	"net/http"

	"github.com/m-zajac/gosnakes/snakes"
)

func NewGame(w http.ResponseWriter, r *http.Request) {
	player1 := snakes.NewRandomPlayer("p1")
	player2 := snakes.NewRandomPlayer("p2")
	board := snakes.Board{
		Size: [2]int{20, 20},
		Snakes: []snakes.Snake{
			snakes.Snake{
				ID:   player1.ID(),
				Head: [2]int{0, 0},
				Tail: [][2]int{},
			},
			snakes.Snake{
				ID:   player2.ID(),
				Head: [2]int{19, 19},
				Tail: [][2]int{},
			},
		},
	}
	g, err := snakes.NewGame(board, player1.ID(), player1, player2)
	if err != nil {
		log.Fatalf("got NewGame err: %v", err)
	}
	_ = g
}
