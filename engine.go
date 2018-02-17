package chesskimo

import (
	"fmt"
	"strings"
)

const (
	ENGINE_STATE_IDLE = iota
	ENGINE_STATE_QUIT
)

type Engine struct {
	// name of the engine including version number.
	name   string
	author string
	// protocol defines how the engine communicates with the frontend.
	protocol Protocol
	// atomicState defines the current state of the engine
	//	atomicState uint32
	//	waitgroup   sync.WaitGroup

	board Board
}

func NewEngine(name, author string, protocol Protocol) *Engine {
	e := &Engine{
		name:     name,
		author:   author,
		protocol: protocol,
		board:    NewBoard(),
	}

	return e
}

func (e *Engine) NewGame() {
	e.board = NewBoard()
}

// Quit shuts everything down gracefully and returns.
func (e *Engine) Quit() {
	//	atomic.StoreUint32(&e.atomicState, ENGINE_STATE_QUIT)
	// TODO wait for search..
}

func (e *Engine) MakeMove(move string) {
	if len(move) >= 4 {
		promo := NONE

		fmt.Println(move[0:2])
		from, err := parseFENSquare(move[0:2])
		if err != nil {
			return
		}

		fmt.Println(move[2:4])
		to, err := parseFENSquare(move[2:4])
		if err != nil {
			return
		}

		if len(move) == 5 {
			switch strings.ToLower(move[4:5]) {
			case "q":
				promo = QUEEN
			case "r":
				promo = ROOK
			case "b":
				promo = BISHOP
			case "n":
				promo = KNIGHT
			}
		}
		from = from.To0x88()
		to = to.To0x88()
		fmt.Println("FROMTO", from, to, promo)
		bm := NewBitMove(from, to, promo)
		fmt.Println("MINI", bm.MiniNotation())
	}

}
