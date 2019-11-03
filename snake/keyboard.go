package snake

import "github.com/nsf/termbox-go"

// keyEvent hold a keyboard event symbol and keyboard key.
type keyEvent struct {
	event rune
	key   termbox.Key
}

// Keyboard events
const (
	RETRY rune = 'r'
	END   rune = 'q'
)

// keyboard hold a keyboard event channel.
type keyboard struct {
	eventChan chan keyEvent
}

// newKeyboard is used to create new keyboard object.
func newKeyboard() *keyboard {
	return &keyboard{eventChan: make(chan keyEvent)}
}

// handleEvent is used to handle keyboard press events.
func (k *keyboard) handleEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft, termbox.KeyArrowRight, termbox.KeyArrowUp, termbox.KeyArrowDown:
				k.eventChan <- keyEvent{key: ev.Key}
			default:
				if ev.Ch == RETRY || ev.Ch == END {
					k.eventChan <- keyEvent{event: ev.Ch, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
