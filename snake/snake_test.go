package snake

import "testing"

func generateSnake() *snake {
	return newSnake([]point{
		{x: 0, y: 2},
		{x: 1, y: 2},
		{x: 2, y: 2},
		{x: 3, y: 2},
		{x: 4, y: 2},
	})
}

func TestSnakeCover(t *testing.T) {
	s := generateSnake()
	ps1 := []point{
		newPoint(0, 2),
		newPoint(3, 2),
	}

	ps2 := []point{
		newPoint(0, 0),
		newPoint(4, 4),
		newPoint(2, 3),
	}

	for _, p := range ps1 {
		if !s.cover(p) {
			t.Errorf("s=%v, s.cover(%v): got false, want true", s, p)
		}
	}

	for _, p := range ps2 {
		if s.cover(p) {
			t.Errorf("s=%v, s.cover(%v): got true, want false", s, p)
		}
	}
}

func TestSnakeHead(t *testing.T) {
	s := generateSnake()
	p := newPoint(4, 2)
	if !s.head().equal(p) {
		t.Errorf("s.head()=%v, s.head().equal(%v): got false, want true", s.head(), p)
	}
}
