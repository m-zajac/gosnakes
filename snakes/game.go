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

type Game struct {
	Arena         [2]int
	Snakes        []Snake
	MovingSnakeID string
	Apples        [][2]int
}

func (g Game) Move(id string, m Move) (Game, error) {
	movingSnakeIndex := -1
	for i, s := range g.Snakes {
		if s.ID == id {
			movingSnakeIndex = i
			break
		}
	}
	if movingSnakeIndex < 0 {
		return g, errors.New("no snake to move")
	}

	s, err := g.Snakes[movingSnakeIndex].Move(m)
	if err != nil {
		return g, fmt.Errorf("invalid move: %v", err)
	}

	// check arena collisions
	if !s.CheckBoundaries(g.Arena[0], g.Arena[1]) {
		return g, errors.New("invalid move")
	}

	// check snakes collisions
	for i := range g.Snakes {
		if i == movingSnakeIndex {
			continue
		}
		if g.Snakes[i].InBody(s.Head) {
			return g, fmt.Errorf("collision with snake %s", g.Snakes[i].ID)
		}
	}

	// check apples
	for i, a := range g.Apples {
		if a[0] == s.Head[0] && a[1] == s.Head[1] {
			// eat!
			s.Tail = make(Tail, len(s.Tail)+1)
			copy(s.Tail[1:], g.Snakes[movingSnakeIndex].Tail)
			s.Tail[0] = g.Snakes[movingSnakeIndex].Head

			// take apple from game
			g.Apples = append(g.Apples[:i], g.Apples[i+1:]...)
		}
	}
	g.Snakes[movingSnakeIndex] = s

	return g, nil
}

func (g Game) String() string {
	arena := make([][]rune, g.Arena[1])
	for i := 0; i < g.Arena[1]; i++ {
		arena[i] = make([]rune, g.Arena[0])
		for j := 0; j < g.Arena[0]; j++ {
			arena[i][j] = EmptyCellRune
		}
	}

	for _, a := range g.Apples {
		arena[a[1]][a[0]] = AppleRune
	}

	for _, s := range g.Snakes {
		arena[s.Head[1]][s.Head[0]] = SnakeHeadRune
		for _, c := range s.Tail {
			arena[c[1]][c[0]] = SnakeTailRune
		}
	}

	var b bytes.Buffer
	b.WriteString("\n┌")
	for i := 0; i < g.Arena[0]; i++ {
		b.WriteString("─")
	}
	b.WriteString("┐\n")

	for i := 0; i < g.Arena[1]; i++ {
		b.WriteString("│")
		for j := 0; j < g.Arena[0]; j++ {
			b.WriteRune(arena[i][j])
		}
		b.WriteString("│\n")
	}

	b.WriteString("└")
	for i := 0; i < g.Arena[0]; i++ {
		b.WriteString("─")
	}
	b.WriteString("┘\n")

	return b.String()
}
