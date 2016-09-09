package snakes

import (
	"errors"
	"math/rand"
)

type Player interface {
	ID() string
	Move(Board) Move
}

type PlayerProvider interface {
	PlayerFromID(id string) (Player, error)
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

type mockPlayerProv struct {
	players map[string]Player
}

func (pp mockPlayerProv) PlayerFromID(id string) (Player, error) {
	player, ok := pp.players[id]
	if !ok {
		return nil, errors.New("invalid player id")
	}
	return player, nil
}

func NewMockPlayerProvider(players ...Player) PlayerProvider {
	pp := mockPlayerProv{
		players: make(map[string]Player, len(players)),
	}
	for _, p := range players {
		pp.players[p.ID()] = p
	}
	return &pp
}
