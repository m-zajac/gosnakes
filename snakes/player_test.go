package snakes

import "testing"

func TestPlayer(t *testing.T) {
	var err error
	board := Board{
		Size: [2]int{2, 1},
		Snakes: []Snake{
			Snake{
				ID:   "p1",
				Head: [2]int{0, 0},
				Tail: [][2]int{},
			},
		},
	}

	player := NewRandomPlayer("p1")

	// go right...
	move := player.Move(board)
	board, err = board.Move(player.ID(), move)
	if err != nil {
		t.Errorf("got board move err: %v", err)
	}
	if board.Snakes[0].Head != [2]int{1, 0} {
		t.Errorf("invalid snake pos. move: %s, snake: %v", string(move), board.Snakes[0].Body())
	}

	// go left...
	move = player.Move(board)
	board, err = board.Move(player.ID(), move)
	if err != nil {
		t.Errorf("got board move err: %v", err)
	}
	if board.Snakes[0].Head != [2]int{0, 0} {
		t.Errorf("invalid snake pos. move: %s, snake: %v", string(move), board.Snakes[0].Body())
	}
}
