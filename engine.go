package chesskimo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ENGINE_STATE_IDLE = iota
	ENGINE_STATE_QUIT
)

var (
	// ErrInvalidMoveNotation that a move is not in the correct notation
	ErrInvalidMoveNotation = errors.New("Move has bad notation formatting")
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

	logger *log.Logger
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

func (e *Engine) Run() {
	f, err := os.Create("chesskimo.log")
	if err != nil {
		panic("Cannot create log file.")
	}
	defer f.Close()
	e.logger = log.New(f, "", log.LstdFlags|log.Lshortfile)

	e.protocol.RunInputOutputLoop(e)
}

func (e *Engine) NewGame() {
	e.board = NewBoard()
}

// Quit shuts everything down gracefully and returns.
func (e *Engine) Quit() {
	//	atomic.StoreUint32(&e.atomicState, ENGINE_STATE_QUIT)
	// TODO wait for search..
}

func (e *Engine) GetLegalMoves() MoveList {
	ml := MoveList{}
	e.board.GenerateAllLegalMoves(&ml)
	return ml
}

func (e *Engine) MakeMove(move string) error {
	if len(move) >= 4 {
		promo := NONE

		fmt.Println(move[0:2])
		from, err := parseFENSquare(move[0:2])
		if err != nil {
			return ErrInvalidMoveNotation
		}

		fmt.Println(move[2:4])
		to, err := parseFENSquare(move[2:4])
		if err != nil {
			return ErrInvalidMoveNotation
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
			default:
				// Impossible promotion.
				return ErrInvalidMoveNotation
			}
		}
		from = from.To0x88()
		to = to.To0x88()
		bm := NewBitMove(from, to, promo)

		e.logger.Print("*** exec move: ", bm.MiniNotation())

		// TODO - test if move is really legal?

		e.board.MakeLegalMove(bm)
		e.logger.Print(e.board.String())
	}

	return nil
}
