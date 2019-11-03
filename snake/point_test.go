package snake

import "testing"

func TestPointEqual(t *testing.T) {
	p1 := newPoint(1, 2)
	p2 := newPoint(3, 4)
	p3 := newPoint(1, 2)
	p4 := newPoint(2, 1)

	if ok := p1.equal(p2); ok {
		t.Errorf("%v equals %v = %t, want false", p1, p2, ok)
	}

	if ok := p1.equal(p4); ok {
		t.Errorf("%v equals %v = %t, want false", p1, p4, ok)
	}

	if ok := p1.equal(p3); !ok {
		t.Errorf("%v equals %v = %t, want true", p1, p3, ok)
	}
}

func TestPointInRectangle(t *testing.T) {
	ps1 := []point{
		newPoint(0, 0),
		newPoint(1, 2),
		newPoint(2, 5),
		newPoint(3, 4),
	}

	ps2 := []point{
		newPoint(7, 6),
		newPoint(1, 7),
		newPoint(2, 9),
		newPoint(10, 0),
	}

	rect := newRectangle(0, 0, 6, 6)
	for _, p := range ps1 {
		if !p.in(rect) {
			t.Errorf("p=%v, p.in(%v): got false, want true", p, rect)
		}
	}

	for _, p := range ps2 {
		if p.in(rect) {
			t.Errorf("p=%v, p.in(%v): got true, want false", p, rect)
		}
	}
}
