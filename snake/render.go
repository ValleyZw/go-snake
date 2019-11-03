package snake

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	POINT = '■' // snake and food symbol
	COL   = '│' // game board column symbol
	ROW   = '─' // game board row symbol
)

// renderSnake is used to draw snake by snake body.
func renderSnake(left, bottom int, s *snake) {
	for _, b := range s.body {
		setCell(left+b.x, bottom-b.y, POINT)
	}
}

// renderFood is used to add food to the playground.
func renderFood(left, bottom int, f point) {
	setCell(left+f.x, bottom-f.y, POINT)
}

// renderBoard is used to initialize a game board by board size.
func renderBoard(width, left, top, right, bottom int) {
	for i := top; i < bottom; i++ {
		setCell(left-1, i, COL)
		setCell(right, i, COL)
	}

	setCell(left-1, top, '┌')
	setCell(left-1, bottom, '└')
	setCell(right, top, '┐')
	setCell(right, bottom, '┘')

	fill(left, top, width, 1, termbox.Cell{Ch: ROW})
	fill(left, bottom, width, 1, termbox.Cell{Ch: ROW})
}

// renderScore is used to show game score.
func renderScore(left, top, s int) {
	printtb(left, top-1, fmt.Sprintf("Score: %d", s))
}

func renderMsg(left, bottom int) {
	printtb(left, bottom+1, fmt.Sprint("r Retry q Quit"))
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func printtb(x, y int, msg string) {
	for _, c := range msg {
		setCell(x, y, c)
		x++
	}
}

func setCell(x, y int, c rune) {
	termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
}
