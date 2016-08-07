package snakes

import "errors"

type Tail [][2]int

func (t Tail) Copy() Tail {
	tc := make(Tail, len(t))
	copy(tc, t)
	return tc
}

type Snake struct {
	ID   string
	Head [2]int
	Tail Tail
}

func (s Snake) Body() [][2]int {
	b := [][2]int{s.Head}
	for _, c := range s.Tail {
		b = append(b, c)
	}

	return b
}

func (s Snake) InTail(c [2]int) bool {
	for _, sc := range s.Tail {
		if c[0] == sc[0] && c[1] == sc[1] {
			return true
		}
	}

	return false
}

func (s Snake) InBody(c [2]int) bool {
	if c[0] == s.Head[0] && c[1] == s.Head[1] {
		return true
	}
	return s.InTail(c)
}

func (s Snake) Move(m Move) (Snake, error) {
	s.Tail = s.Tail.Copy() // new tail slice
	newHead := moveCell(s.Head, m)
	if s.InTail(newHead) {
		return s, errors.New("self collision")
	}

	if len(s.Tail) > 0 {
		copy(s.Tail[1:], s.Tail)
		s.Tail[0] = s.Head
	}
	s.Head = newHead

	return s, nil
}

func (s Snake) CheckBoundaries(arenaWidth, arenaHeight int) bool {
	for _, c := range s.Body() {
		if !checkCellBoundaries(c, arenaWidth, arenaHeight) {
			return false
		}
	}
	return true
}

func moveCell(c [2]int, m Move) [2]int {
	switch m {
	case 'u':
		c[1]--
	case 'd':
		c[1]++
	case 'l':
		c[0]--
	case 'r':
		c[0]++
	}
	return c
}

func checkCellBoundaries(c [2]int, arenaWidth int, arenaHeight int) bool {
	if c[0] < 0 || c[0] >= arenaWidth {
		return false
	}
	if c[1] < 0 || c[1] >= arenaHeight {
		return false
	}
	return true
}
