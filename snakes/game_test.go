package snakes

import "testing"

func TestCollisions(t *testing.T) {
	var err error
	g := Game{
		Arena: [2]int{4, 4},
		Snakes: []Snake{
			Snake{
				ID:   "1",
				Head: [2]int{0, 0},
				Tail: [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 3}},
			},
			Snake{
				ID:   "2",
				Head: [2]int{1, 2},
				Tail: [][2]int{{1, 1}},
			},
		},
	}

	testMoves := []struct {
		snakeID     string
		move        Move
		expectError bool
	}{
		// wall - top, left
		{"1", MoveUp, true},
		{"1", MoveLeft, true},

		{"1", MoveRight, false},
		{"1", MoveRight, false},
		{"1", MoveRight, false},

		// wall - right
		{"1", MoveRight, true},

		{"1", MoveDown, false},
		{"1", MoveDown, false},
		{"1", MoveDown, false},

		// wall - bottom
		{"1", MoveDown, true},

		// self collision
		{"1", MoveDown, true},

		{"1", MoveLeft, false},
		{"1", MoveUp, false},

		// self collision with tail
		{"1", MoveRight, true},

		// collision with s2 head
		{"1", MoveLeft, true},

		{"1", MoveUp, false},

		// collision with s2 tail
		{"1", MoveLeft, true},
	}

	t.Log(g.String())
	for i, tm := range testMoves {
		t.Logf("moving %s %s...", tm.snakeID, string(tm.move))
		g, err = g.Move(tm.snakeID, tm.move)
		if err == nil && tm.expectError {
			t.Fatalf("move %d - expeted error", i)
		} else if err != nil && !tm.expectError {
			t.Fatalf("move %d - unexpeted error: %v", i, err)
		}
		t.Log(g.String())
	}
}

func TestApples(t *testing.T) {
	var err error
	g := Game{
		Arena: [2]int{3, 4},
		Snakes: []Snake{
			Snake{
				ID:   "1",
				Head: [2]int{0, 0},
				Tail: [][2]int{},
			},
		},
		Apples: [][2]int{{0, 1}, {1, 1}},
	}

	testMoves := []struct {
		snakeID string
		move    Move
		s1Size  int
		apples  int
	}{
		{"1", MoveRight, 1, 2},
		{"1", MoveRight, 1, 2},
		{"1", MoveDown, 1, 2},

		// eat
		{"1", MoveLeft, 2, 1},
		{"1", MoveLeft, 3, 0},

		// afterparty
		{"1", MoveDown, 3, 0},
	}

	t.Log(g.String())
	for i, tm := range testMoves {
		t.Logf("moving %s %s...", tm.snakeID, string(tm.move))
		g, err = g.Move(tm.snakeID, tm.move)
		if err != nil {
			t.Fatalf("move %d - unexpeted error: %v", i, err)
		}
		if len(g.Apples) != tm.apples {
			t.Fatalf("move %d - unexpeted num of apples, want %d, got %d", i, tm.apples, len(g.Apples))
		}
		t.Log(g.String())
	}
}
