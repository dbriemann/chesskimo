package chesskimo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type UCI struct {
	newGame bool
}

func (u *UCI) RunInputOutputLoop(engine *Engine) {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		panic("Cannot read from std in.")
	}

	for {
		command, err := reader.ReadString('\n')
		engine.logger.Print("<-- cmd: ", command)
		if err == io.EOF {
			break
		} else if err == nil && len(command) > 0 {
			// Split input string into command parts.
			input := strings.Split(strings.Trim(command, " \t\r\n"), " ")
			cmd := input[0]
			switch cmd {
			case "quit":
				// TODO -> shutdown engine.
				os.Exit(0)
			case "uci":
				// Enable UCI mode and identify yourself.
				u.cmdUci(engine)
			case "isready":
				// Check if engine can receive commands/is active.
				u.cmdIsready()
			case "ucinewgame":
				u.cmdNewGame(engine)
			case "position":
				u.cmdPosition(engine, input[1:])
			case "go":
				u.cmdGo(engine, input[1:])
			case "stop":
				u.cmdStop(engine)
			}
		}
	}

	// TODO -> shutdown engine.
	os.Exit(0)
}

func (u *UCI) cmdStop(engine *Engine) {
	// TODO stop engine
}

func (u *UCI) cmdGo(engine *Engine, args []string) {
	for len(args) > 0 {
		cmd := args[0]
		args = args[1:]
		switch cmd {
		case "searchmoves":
		case "ponder":
		case "wtime":
		case "btime":
		case "winc":
		case "binc":
		case "movestogo":
		case "depth":
		case "nodes":
		case "mate":
		case "movetime":
		case "infinite":
		}
	}

	//	moves := engine.GetLegalMoves()
	//	r := rand.Intn(int(moves.Size))
	//	bm := moves.Moves[r]
	ts := TimeSettings{}
	dostop := uint32(0)
	sr := engine.search(engine, &ts, &dostop)
	engine.logger.Println("--> best move:", sr.BestMove.MiniNotation())
	engine.board.MakeLegalMove(sr.BestMove)
	engine.logger.Print(engine.board.String())
	fmt.Println("bestmove", sr.BestMove.MiniNotation())
}

func (u *UCI) cmdPosition(engine *Engine, args []string) {
	if len(args) > 1 {
		if !u.newGame {
			// Continue ongoing game by just extracting the last move.
			engine.logger.Print("*** continue ongoing game")
			last := args[len(args)-1]
			err := engine.MakeMove(last)
			if err != nil {
				// UCI is crap.
				engine.logger.Print("*** move ", last, " impossible: ", err.Error())
			}
		} else {
			first := args[0]
			if first == "startpos" {
				args = args[1:]
				engine.board.SetStartingPosition()
			} else {
				if first == "fen" {
					// Some frontends say "position fen" then specify the actual fen, some specify
					// the actual fen directly after "position", so we have to check both ways..
					// no wonder with this as official doc: http://wbec-ridderkerk.nl/html/UCIProtocol.html
					// *sigh*
					args = args[1:]
					if len(args) > 0 {
						first = args[0]
					} else {
						return // Invalid.. missing arguments
					}
				}

				fen := ""
				// We now have to concat all FEN parts.
				for i := 0; i < len(args); i++ {
					if args[i] != "moves" {
						// TODO - remove all quotes... just to be sure.
						fen += args[i] + " "
					} else {
						// Finished FEN.. discard fen arguments.
						args = args[i:]
					}
				}
				err := engine.board.SetFEN(fen)
				if err != nil {
					fmt.Println("--> error: ", err.Error())
					return // UCI ignores bad commands.
				}
			}

			// Last check for "moves" subcommand.
			// Must at least have length 2 (including a move).
			if len(args) > 1 && args[0] == "moves" {
				moves := args[1:]
				for _, m := range moves {
					err := engine.MakeMove(m)
					if err != nil {
						// UCI is crap.
						engine.logger.Print("*** move ", m, " impossible: ", err.Error())
					}
				}
			}
		}
		u.newGame = false
	}
}

func (u *UCI) cmdNewGame(engine *Engine) {
	u.newGame = true
	engine.NewGame()
}

func (u *UCI) cmdIsready() {
	// TODO wait for engine ?
	fmt.Println("readyok")
}

func (u *UCI) cmdUci(engine *Engine) {
	fmt.Println("id name", engine.name)
	fmt.Println("id author", engine.author)
	// TODO -> add all possible options here.
	fmt.Println("uciok")
}
