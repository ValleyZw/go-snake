package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

// Game holds a game's game board and keyboard.
type Game struct {
	board    *board
	keyboard *keyboard
}

const (
	snakeLength     = 4   // initial snake length
	refreshInterval = 100 // game refresh interval in ms
)

// NewGame is used to create new Game object.
func NewGame() *Game {
	return &Game{board: initBoard(), keyboard: newKeyboard()}
}

// Start starts the game.
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	go g.keyboard.handleEvent()
	g.board.render(true)
LOOP:
	for {
		select {
		case e := <-g.keyboard.eventChan:
			switch e.event {
			case END:
				break LOOP
			case RETRY:
				g.retry()
			default:
				g.board.snake.changeDirection(e.key)
			}
		default:
			if ok := g.board.snakeMove(); !ok {
				break
			}

			g.board.render(false)
			time.Sleep(g.interval())
		}
	}
}

// initSnake is used to create new snake with default length.
func initSnake() *snake {
	body := make([]point, snakeLength)
	for i := range body {
		body[i] = newPoint(i+1, 1)
	}
	return newSnake(body)
}

// initBoard is used to create new game board object.
func initBoard() *board {
	return newBoard(initSnake())
}

// getScore is used to get game score.
func (g *Game) getScore() int {
	return g.board.score
}

// interval returns the sleep duration of the game in millisecond.
func (g *Game) interval() time.Duration {
	ms := refreshInterval - g.getScore()
	if ms < 1 {
		ms = 1
	}
	return time.Duration(ms) * time.Millisecond
}

// retry is used to refresh the game with default settings.
func (g *Game) retry() {
	g.board = initBoard()
	g.board.render(true)
}
