package snake

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

// board holds a game board's area, food, current game score and snake.
type board struct {
	area  rectangle
	food  point
	score int
	snake *snake
}

const (
	widthMargin, heightMargin = 10, 5
)

// newBoard is used to create new board object.
func newBoard(s *snake) *board {
	rand.Seed(time.Now().UnixNano())
	return &board{snake: s}
}

// snakeMove is used to move a snake of a board.
func (b *board) snakeMove() bool {
	if ok := b.snake.move(); !ok {
		return ok
	}

	if b.snakeOut() {
		return false
	}

	if b.snakeEat() {
		b.score++
		b.snake.length++
		b.addFood()
	}
	return true
}

// snakeOut reports whether a snake left current board area.
func (b *board) snakeOut() bool {
	return !b.snake.head().in(b.area)
}

func (b *board) addFood() {
	var x, y int
	for {
		x, y = rand.Intn(b.area.max.x), rand.Intn(b.area.max.y)
		if !b.snake.cover(newPoint(x, y)) {
			break
		}
	}
	b.food = newPoint(x, y)
}

// snakeEat reports whether snake eat a food.
func (b *board) snakeEat() bool {
	return b.food.equal(b.snake.head())
}

// render is used to init and refresh game ui.
//noinspection ALL
func (b *board) render(init bool) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	termWidth, termHeight := termbox.Size()
	if termWidth < 1 || termHeight < 1 {
		termWidth, termHeight = widthMargin*2, heightMargin*2
	}

	midHeight := termHeight / 2
	boardWidth, boardHeight := termWidth-widthMargin, termHeight-heightMargin
	left := (termWidth - boardWidth) / 2
	right := (termWidth + boardWidth) / 2
	top := midHeight - boardHeight/2
	bottom := midHeight + boardHeight/2
	b.area = newRectangle(0, 0, boardWidth, boardHeight)
	if init {
		b.addFood()
	}

	renderBoard(boardWidth, left, top, right, bottom)
	renderSnake(left, bottom, b.snake)
	renderFood(left, bottom, b.food)
	renderScore(left, top, b.score)
	renderMsg(left, bottom)
}
