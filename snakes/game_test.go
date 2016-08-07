package snakes

import "testing"

func TestGame(t *testing.T) {
	player := NewRandomPlayer("p1")
	board := Board{
		Size: [2]int{2, 1},
		Snakes: []Snake{
			Snake{
				ID:   player.ID(),
				Head: [2]int{0, 0},
				Tail: [][2]int{},
			},
		},
	}
	game, err := NewGame(board, player.ID(), player)
	if err != nil {
		t.Fatalf("got NewGame err: %v", err)
	}

	game.Next()
}
