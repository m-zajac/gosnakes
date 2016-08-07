package snakes

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	EmptyCellRune rune = ' '
	AppleRune     rune = 'a'
	SnakeHeadRune rune = 'S'
	SnakeTailRune rune = 's'

	MoveUp    Move = 'u'
	MoveDown  Move = 'd'
	MoveLeft  Move = 'l'
	MoveRight Move = 'r'
)

type Move rune

type Board struct {
	Size   [2]int
	Snakes []Snake
	Apples [][2]int
}

func (board Board) Move(id string, m Move) (Board, error) {
	movingSnakeIndex := -1
	for i, s := range board.Snakes {
		if s.ID == id {
			movingSnakeIndex = i
			break
		}
	}
	if movingSnakeIndex < 0 {
		return board, errors.New("no snake to move")
	}

	s, err := board.Snakes[movingSnakeIndex].Move(m)
	if err != nil {
		return board, fmt.Errorf("invalid move: %v", err)
	}

	// check arena collisions
	if !s.CheckBoundaries(board.Size[0], board.Size[1]) {
		return board, fmt.Errorf("invalid move: %s", string(m))
	}

	// check snakes collisions
	for i := range board.Snakes {
		if i == movingSnakeIndex {
			continue
		}
		if board.Snakes[i].InBody(s.Head) {
			return board, fmt.Errorf("collision with snake %s", board.Snakes[i].ID)
		}
	}

	// check apples
	for i, a := range board.Apples {
		if a[0] == s.Head[0] && a[1] == s.Head[1] {
			// eat!
			s.Tail = make(Tail, len(s.Tail)+1)
			copy(s.Tail[1:], board.Snakes[movingSnakeIndex].Tail)
			s.Tail[0] = board.Snakes[movingSnakeIndex].Head

			// take apple from game
			apples := make([][2]int, len(board.Apples))
			copy(apples, board.Apples)
			apples = append(apples[:i], apples[i+1:]...)
			board.Apples = apples

			break
		}
	}

	// update new board snakes
	snakes := make([]Snake, len(board.Snakes))
	copy(snakes, board.Snakes)
	snakes[movingSnakeIndex] = s
	board.Snakes = snakes

	// board.Snakes[movingSnakeIndex] = s

	return board, nil
}

func (board Board) String() string {
	arena := make([][]rune, board.Size[1])
	for i := 0; i < board.Size[1]; i++ {
		arena[i] = make([]rune, board.Size[0])
		for j := 0; j < board.Size[0]; j++ {
			arena[i][j] = EmptyCellRune
		}
	}

	for _, a := range board.Apples {
		arena[a[1]][a[0]] = AppleRune
	}

	for _, s := range board.Snakes {
		arena[s.Head[1]][s.Head[0]] = SnakeHeadRune
		for _, c := range s.Tail {
			arena[c[1]][c[0]] = SnakeTailRune
		}
	}

	var b bytes.Buffer
	b.WriteString("\n┌")
	for i := 0; i < board.Size[0]; i++ {
		b.WriteString("─")
	}
	b.WriteString("┐\n")

	for i := 0; i < board.Size[1]; i++ {
		b.WriteString("│")
		for j := 0; j < board.Size[0]; j++ {
			b.WriteRune(arena[i][j])
		}
		b.WriteString("│\n")
	}

	b.WriteString("└")
	for i := 0; i < board.Size[0]; i++ {
		b.WriteString("─")
	}
	b.WriteString("┘\n")

	return b.String()
}
