package snakes

import (
	"errors"
	"fmt"
)

type Game interface {
	Next() error
	Board() Board
}

type game struct {
	players           []Player
	board             Board
	activePlayerIndex int
}

func NewGame(board Board, activePlayerID string, players ...Player) (Game, error) {
	playerIndex := -1
	for i, p := range players {
		if p.ID() == activePlayerID {
			playerIndex = i
			break
		}
	}
	if playerIndex < 0 {
		return nil, errors.New("invalid active player id")
	}
	return &game{
		players:           players,
		board:             board,
		activePlayerIndex: playerIndex,
	}, nil
}

func (g *game) Next() error {
	if len(g.players) == 0 {
		return errors.New("no players")
	}

	player := g.players[g.activePlayerIndex]
	move := player.Move(g.board)
	board, err := g.board.Move(player.ID(), move)
	if err != nil {
		return fmt.Errorf("player %s move err: %v", player.ID(), err)
	}

	g.board = board
	g.activePlayerIndex = (g.activePlayerIndex + 1) % len(g.players)

	return nil
}

func (g game) Board() Board {
	return g.board
}
