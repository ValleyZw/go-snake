package snake

import "github.com/nsf/termbox-go"

// snake holds a snake's body, body length and lt's moving direction.
type snake struct {
	body      []point
	direction termbox.Key
	length    int
}

var opposites = map[termbox.Key]termbox.Key{
	termbox.KeyArrowRight: termbox.KeyArrowLeft,
	termbox.KeyArrowLeft:  termbox.KeyArrowRight,
	termbox.KeyArrowUp:    termbox.KeyArrowDown,
	termbox.KeyArrowDown:  termbox.KeyArrowUp,
}

// newSnake is used to create new snake object.
func newSnake(b []point) *snake {
	return &snake{
		length:    len(b),
		body:      b,
		direction: termbox.KeyArrowRight,
	}
}

func (s *snake) changeDirection(d termbox.Key) {
	if o := opposites[d]; o != s.direction {
		s.direction = d
	}
}

func (s *snake) head() point {
	return s.body[len(s.body)-1]
}

func (s *snake) move() bool {
	p := newPoint(s.head().x, s.head().y)

	switch s.direction {
	case termbox.KeyArrowRight:
		p.x++
	case termbox.KeyArrowLeft:
		p.x--
	case termbox.KeyArrowUp:
		p.y++
	case termbox.KeyArrowDown:
		p.y--
	}

	if s.cover(p) {
		return false
	}

	if s.length > len(s.body) {
		s.body = append(s.body, p)
	} else {
		s.body = append(s.body[1:], p)
	}
	return true
}

// cover determines whether a given point is covered by snake body.
func (s *snake) cover(p point) bool {
	for _, b := range s.body {
		if b.equal(p) {
			return true
		}
	}
	return false
}
