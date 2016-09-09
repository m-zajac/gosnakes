package snakes

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Game interface {
	Next() error
	Board() Board

	encoding.TextMarshaler
}

type game struct {
	ID                string   `json:"id"`
	Players           []Player `json:"-"`
	GameBoard         Board    `json:"board"`
	ActivePlayerIndex int      `json:"activePlayerIndex"`
}

// for json marshall - raw struct without encoding.TextMarshaler
type jsonGame game

type marshaledGame struct {
	jsonGame
	PlayerIDs []string `json:"players"`
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
		ID:                uuid.NewV4().String(),
		Players:           players,
		GameBoard:         board,
		ActivePlayerIndex: playerIndex,
	}, nil
}

func (g *game) Next() error {
	if len(g.Players) == 0 {
		return errors.New("no players")
	}

	player := g.Players[g.ActivePlayerIndex]
	move := player.Move(g.GameBoard)
	board, err := g.GameBoard.Move(player.ID(), move)
	if err != nil {
		return fmt.Errorf("player %s move err: %v", player.ID(), err)
	}

	g.GameBoard = board
	g.ActivePlayerIndex = (g.ActivePlayerIndex + 1) % len(g.Players)

	return nil
}

func (g game) Board() Board {
	return g.GameBoard
}

func (g game) MarshalText() (text []byte, err error) {
	mg := marshaledGame{jsonGame: jsonGame(g)}
	for _, p := range g.Players {
		mg.PlayerIDs = append(mg.PlayerIDs, p.ID())
	}
	data, err := json.Marshal(mg)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %v", err)
	}

	return data, nil
}

func UnmarshalGame(data []byte, playerProvider PlayerProvider) (Game, error) {
	mg := marshaledGame{}
	err := json.Unmarshal(data, &mg)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err: %v", err)
	}

	for _, id := range mg.PlayerIDs {
		player, err := playerProvider.PlayerFromID(id)
		if err != nil {
			return nil, fmt.Errorf("player %s not found", id)
		}
		mg.jsonGame.Players = append(mg.jsonGame.Players, player)
	}

	g := game(mg.jsonGame)
	return &g, nil
}
