package snakes

import (
	"bytes"
	"testing"
)

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
	g, err := NewGame(board, player.ID(), player)
	if err != nil {
		t.Fatalf("got NewGame err: %v", err)
	}

	if id := g.(*game).ID; id == "" {
		t.Error("empty game id")
	} else {
		t.Logf("game id: %s", id)
	}

	if err = g.Next(); err != nil {
		t.Fatalf("game Next err: %v", err)
	}
}

func TestGameMarshal(t *testing.T) {
	player1 := NewRandomPlayer("p1")
	player2 := NewRandomPlayer("p2")
	playerProvider := NewMockPlayerProvider(player1, player2)

	board := Board{
		Size: [2]int{10, 10},
		Snakes: []Snake{
			Snake{
				ID:   player1.ID(),
				Head: [2]int{0, 0},
				Tail: [][2]int{},
			},
			Snake{
				ID:   player2.ID(),
				Head: [2]int{9, 9},
				Tail: [][2]int{},
			},
		},
		Apples: [][2]int{{5, 5}, {7, 7}},
	}
	g, err := NewGame(board, player1.ID(), player1, player2)
	if err != nil {
		t.Fatalf("got NewGame err: %v", err)
	}

	data, err := g.MarshalText()
	if err != nil {
		t.Fatalf("game marshal err: %v", err)
	}
	t.Logf("game: %s", string(data))

	newG, err := UnmarshalGame(data, playerProvider)
	if err != nil {
		t.Fatalf("game unmarshal err: %v", err)
	}

	data2, err := newG.MarshalText()
	if err != nil {
		t.Fatalf("unmarshaled game marshal err: %v", err)
	}
	t.Logf("game marshaled second time: %s", string(data2))

	if !bytes.Equal(data, data2) {
		t.Errorf("second marshal results in different data")
	}
}
