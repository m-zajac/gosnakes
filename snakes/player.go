package snakes

import "math/rand"

type Player interface {
	ID() string
	Move(Board) Move
}

type randomPlayer struct {
	id string
}

func NewRandomPlayer(id string) Player {
	return randomPlayer{
		id: id,
	}
}

func (p randomPlayer) ID() string {
	return p.id
}

func (p randomPlayer) Move(b Board) Move {
	moves := []Move{MoveDown, MoveUp, MoveLeft, MoveRight}
	for _, i := range rand.Perm(len(moves)) {
		if _, err := b.Move(p.ID(), moves[i]); err == nil {
			return moves[i]
		}
	}
	return moves[0]
}
