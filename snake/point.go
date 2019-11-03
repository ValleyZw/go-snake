package snake

// A point is an X, Y coordinate pair. The axes increase right and down.
type point struct {
	x, y int
}

// newPoint is used to create a new point object.
func newPoint(x, y int) point {
	return point{x, y}
}

// equal reports whether p and q are equal.
func (p point) equal(q point) bool {
	return p == q
}

// in reports whether p is in r.
func (p point) in(r rectangle) bool {
	return r.min.x <= p.x && p.x <= r.max.x && r.min.y <= p.y && p.y <= r.max.y
}

// A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
type rectangle struct {
	min, max point
}

// newRectangle is used to create a new rectangle object.
func newRectangle(left, top, right, bottom int) rectangle {
	if left > right {
		left, right = right, left
	}
	if top > bottom {
		top, bottom = bottom, top
	}
	return rectangle{min: point{left, top}, max: point{right, bottom}}
}
