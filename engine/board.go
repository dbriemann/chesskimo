package engine

import (
	"fmt"
	"strconv"

	"github.com/dbriemann/chesskimo/base"
	"github.com/dbriemann/chesskimo/fen"
)

// Board contains all information for a chess board state, including the board itself as 0x88 board.
//
//   +------------------------+
// 8 |70 71 72 73 74 75 76 77 | 78 79 7a 7b 7c 7d 7e 7f
// 7 |60 61 62 63 64 65 66 67 | 68 69 6a 6b 6c 6d 6e 6f
// 6 |50 51 52 53 54 55 56 57 | 58 59 5a 5b 5c 5d 5e 5f
// 5 |40 41 42 43 44 45 46 47 | 48 49 4a 4b 4c 4d 4e 4f
// 4 |30 31 32 33 34 35 36 37 | 38 39 3a 3b 3c 3d 3e 3f
// 3 |20 21 22 23 24 25 26 27 | 28 29 2a 2b 2c 2d 2e 2f
// 2 |10 11 12 13 14 15 16 17 | 18 19 1a 1b 1c 1d 1e 1f
// 1 | 0  1  2  3  4  5  6  7 |  8  9  a  b  c  d  e  f
//   +------------------------+
//     a  b  c  d  e  f  g  h
//
// All indexes shown are in HEX. The left board represents the actual playing board, whereas there right board
// is for detection of illegal moves without heavy conditional usage.
type Board struct {
	Squares     [64 * 2]base.Piece
	CastleShort [2]bool
	CastleLong  [2]bool
	MoveNumber  uint16
	DrawCounter uint16
	CheckInfo   base.Square
	EpSquare    base.Square
	Player      base.Color
	Kings       [2]base.Square
	Queens      [2]base.PieceList
	Rooks       [2]base.PieceList
	Bishops     [2]base.PieceList
	Knights     [2]base.PieceList
	Pawns       [2]base.PieceList
}

const (
	CHECK_NONE         = base.OTB
	CHECK_DOUBLE_CHECK = 0x0F // Is also off the board but used as double check value.
	CHECK_CHECKMATE    = 0x1F
)

var (
	// Lookup0x88 maps the indexes of a 8x8 MinBoard to the 0x88 board indexes.
	Lookup0x88 = [64]base.Square{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
		0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
	}
	SQUARE_DIFFS = [240]base.Piece{}
	DIFF_DIRS    = [240]int8{}
)

func init() {
	populateDiffMaps()
}

func populateDiffMaps() {
	diffdir := int8(0)
	for _, from := range Lookup0x88 {
		for _, to := range Lookup0x88 {
			diff := uint8(0x77 + (int8(from) - int8(to)))
			pbits := base.Piece(0) // Resulting bitset.

			// Test if diff is a possible move for all types of pieces.
			// kings:
			for _, dir := range base.KING_DIRS {
				target := base.Square(int8(from) + dir)

				if target == to {
					pbits |= base.KING
					break
				}
			}

			// diagonal sliders:
			for _, dir := range base.DIAGONAL_DIRS {
				for steps := int8(1); ; steps++ {
					target := base.Square(int8(from) + steps*dir)
					if !target.OnBoard() {
						// We left the board -> next dir.
						break
					}

					if target == to {
						diffdir = dir
						// Can be reachedby queens and bishops (diagonally).
						pbits |= (base.BISHOP | base.QUEEN)
						// The goto acts as a double break.
						// Skipping orthos and knights.
						goto INSPECTION_FINISHED
					}
				}
			}

			// orthogonal sliders:
			for _, dir := range base.ORTHOGONAL_DIRS {
				for steps := int8(1); ; steps++ {
					target := base.Square(int8(from) + steps*dir)
					if !target.OnBoard() {
						// We left the board -> next dir.
						break
					}

					if target == to {
						diffdir = dir
						// Can be reachedby queens and bishops (diagonally).
						pbits |= (base.ROOK | base.QUEEN)
						// The goto acts as a double break.
						// Skipping knights.
						goto INSPECTION_FINISHED
					}
				}
			}

			// knights:
			for _, dir := range base.KNIGHT_DIRS {
				target := base.Square(int8(from) + dir)

				if target == to {
					diffdir = dir
					pbits |= base.KNIGHT
					break
				}
			}

		INSPECTION_FINISHED:

			DIFF_DIRS[diff] = diffdir
			SQUARE_DIFFS[diff] = pbits
		}
	}
}

func NewBoard() Board {
	b := Board{}
	b.Kings = [2]base.Square{base.OTB, base.OTB}
	b.Queens = [2]base.PieceList{base.NewPieceList()}
	b.Rooks = [2]base.PieceList{base.NewPieceList()}
	b.Bishops = [2]base.PieceList{base.NewPieceList()}
	b.Knights = [2]base.PieceList{base.NewPieceList()}
	b.Pawns = [2]base.PieceList{base.NewPieceList()}

	b.SetStartingPosition()
	return b
}

// SetStartingPosition sets the starting position of a default chess game via FEN code.
func (b *Board) SetStartingPosition() {
	err := b.SetFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
}

// SetFEN sets a chess position from a 'FEN' string. This function is expensive
// in terms of computation time. It should only be used in non performance critical
// code such as setting up states etc.
func (b *Board) SetFEN(fenstr string) error {
	mb, err := fen.ParseFEN(fenstr)
	if err != nil {
		panic(err)
	}

	for col := base.BLACK; col <= base.WHITE; col++ {
		b.Queens[col].Clear()
		b.Rooks[col].Clear()
		b.Bishops[col].Clear()
		b.Knights[col].Clear()
		b.Pawns[col].Clear()
	}

	// No error encountered -> set all members of the board instance.
	b.MoveNumber = mb.MoveNum
	b.DrawCounter = mb.HalfMoves

	if mb.EpSquare != base.OTB {
		b.EpSquare = Lookup0x88[mb.EpSquare]
	} else {
		b.EpSquare = base.OTB
	}

	b.CastleShort = mb.CastleShort
	b.CastleLong = mb.CastleLong
	b.Player = mb.Color
	for idx, piece := range mb.Squares {
		sq := Lookup0x88[idx]
		b.Squares[sq] = piece
		switch piece {
		case base.WKING:
			b.Kings[base.WHITE] = sq
		case base.BKING:
			b.Kings[base.BLACK] = sq
		case base.WQUEEN:
			b.Queens[base.WHITE].Add(sq)
		case base.BQUEEN:
			b.Queens[base.BLACK].Add(sq)
		case base.WROOK:
			b.Rooks[base.WHITE].Add(sq)
		case base.BROOK:
			b.Rooks[base.BLACK].Add(sq)
		case base.WBISHOP:
			b.Bishops[base.WHITE].Add(sq)
		case base.BBISHOP:
			b.Bishops[base.BLACK].Add(sq)
		case base.WKNIGHT:
			b.Knights[base.WHITE].Add(sq)
		case base.BKNIGHT:
			b.Knights[base.BLACK].Add(sq)
		case base.WPAWN:
			b.Pawns[base.WHITE].Add(sq)
		case base.BPAWN:
			b.Pawns[base.BLACK].Add(sq)
		}
	}

	// Set info board and find possible checks.
	//	b.BuilInfoBoard(b.Player)

	return nil
}

func (b *Board) clearInfoBoard() {
	for _, idx := range base.INFO_BOARD_INDEXES {
		b.Squares[idx] = base.NO_PIECE
	}
}

func (b *Board) DetectChecksAndPins(color base.Color) {
	b.clearInfoBoard()
	oppColor := color.Flip()

	kingSq := b.Kings[color]
	checkCounter := 0

	//	fmt.Println("KING:", base.PrintBoardIndex[kingSq])

	// 1. Detect checks by knights.
	for i := uint8(0); i < b.Knights[oppColor].Size; i++ {
		knightSq := b.Knights[oppColor].Pieces[i]
		diff := knightSq.Diff(kingSq)
		if SQUARE_DIFFS[diff].Contains(base.KNIGHT) {
			// The knight checks the king. Save the info and break the loop.
			// There can never be a double check of 2 knights.
			b.CheckInfo = knightSq
			checkCounter++
			b.Squares[knightSq.ToInfoIndex()] = base.INFO_CHECK
			break
		}
	}

	// 2. Detect checks by pawns.
	for i := uint8(0); i < b.Pawns[oppColor].Size; i++ {
		pawnSq := b.Pawns[oppColor].Pieces[i]
		for _, dir := range base.PAWN_CAPTURE_DIRS[oppColor] {
			to := base.Square(int8(pawnSq) + dir)
			if kingSq == to {
				// The pawn checks the king. Save the info and break the loop.
				// There can never be a double check of 2 pawns.
				b.CheckInfo = pawnSq
				checkCounter++
				b.Squares[pawnSq.ToInfoIndex()] = base.INFO_CHECK
				goto EXIT_PAWN_CHECK // Acts as double break.. nothing harmful really... :)
			}
		}
	}

EXIT_PAWN_CHECK:

	// 2.b No double check testing is needed. Pawn and knight cannot give double check.

	// 3. Detect checks (and pins) by Sliders.
	// 3.a) bishops
	checkCounter += b.DetectSliderChecksAndPins(color, checkCounter, &b.Bishops[oppColor], &base.DIAGONAL_DIRS, base.BISHOP)
	if checkCounter > 1 {
		// Double check detected -> everything else is not of interest now.
		b.CheckInfo = CHECK_DOUBLE_CHECK
		goto PRETTY_PRINT_RETURN
	}
	// 3.b) queens diagonally
	checkCounter += b.DetectSliderChecksAndPins(color, checkCounter, &b.Queens[oppColor], &base.DIAGONAL_DIRS, base.QUEEN)
	if checkCounter > 1 {
		// Double check detected -> everything else is not of interest now.
		b.CheckInfo = CHECK_DOUBLE_CHECK
		goto PRETTY_PRINT_RETURN
	}
	// 3.c) rooks
	checkCounter += b.DetectSliderChecksAndPins(color, checkCounter, &b.Rooks[oppColor], &base.ORTHOGONAL_DIRS, base.ROOK)
	if checkCounter > 1 {
		// Double check detected -> everything else is not of interest now.
		b.CheckInfo = CHECK_DOUBLE_CHECK
		goto PRETTY_PRINT_RETURN
	}
	// 3.d) queens orthogonally
	checkCounter += b.DetectSliderChecksAndPins(color, checkCounter, &b.Queens[oppColor], &base.ORTHOGONAL_DIRS, base.ROOK)
	if checkCounter > 1 {
		// Double check detected -> everything else is not of interest now.
		b.CheckInfo = CHECK_DOUBLE_CHECK
		goto PRETTY_PRINT_RETURN
	}

PRETTY_PRINT_RETURN:
	fmt.Println(b.InfoBoardString())
	//	for _, idx := range base.INFO_BOARD_INDEXES {
	//		fmt.Print(int(b.Squares[idx]))
	//	}
	fmt.Println(b.CheckInfo)
	fmt.Println()
}

func (b *Board) DetectSliderChecksAndPins(color base.Color, ccount int, pieceList *base.PieceList, dirList *[4]int8, ptype base.Piece) int {
	kingSq := b.Kings[color]
	checkCounter := 0

	for i := uint8(0); i < pieceList.Size; i++ {
		sliderSq := pieceList.Pieces[i]
		//		fmt.Println("SLIDER", base.PrintBoardIndex[sliderSq])
		diff := kingSq.Diff(sliderSq)
		if SQUARE_DIFFS[diff].Contains(ptype) {
			// The slider possibly checks the king or pins a piece.
			diffdir := DIFF_DIRS[diff]
			//			fmt.Println("POSSIBLE PATH DIR:", diffdir)
			// Starting from the king we step through the diagonal in question towards the enemy slider.
			info := base.INFO_NONE
			for stepSq := base.Square(int8(kingSq) + diffdir); ; stepSq = base.Square(int8(stepSq) + diffdir) {
				//				fmt.Println("PATH SQ:", stepSq, base.PrintBoardIndex[stepSq])
				curPiece := b.Squares[stepSq]
				if curPiece.IsEmpty() {
					//					fmt.Println("EMTPY -> CONTINUE")
					continue
				} else if curPiece.HasColor(color) {
					// A friendly piece was encountered
					if info == base.INFO_NONE {
						//						fmt.Println("FRIENDLY -> SET ATTACKED")
						// This is the first piece on the path -> mark it as attacked.
						info = base.INFO_ATTACKED
					} else {
						//						fmt.Println("FRIENDLY -> REMOVE ATTACK")
						// We already marked a piece on this path before -> there cannot be a pin anymore.
						// (There is an exception for EP capture but those are handled elsewhere.
						info = base.INFO_NONE
						break // Nothing more to do on this path.
					}
				} else {
					// An enemy piece was encountered
					if curPiece.Contains(ptype) {
						//						fmt.Println("ENEMY -> SLIDER")
						// We reached the checker/pinner -> decide.
						if info == base.INFO_NONE {
							//							fmt.Println("CHECK")
							// It's a check!
							checkCounter++
							b.CheckInfo = sliderSq
							info = base.INFO_CHECK
							break
						} else {
							//							fmt.Println("PIN")
							// It's a pinner!
							info = base.INFO_PIN
							break
						}
					} else {
						//						fmt.Println("ENEMY -> OTHER PIECE")
						// The enemy piece cannot pin or check. Break out.
						info = base.INFO_NONE
						break
					}
				}
			}

			if info == base.INFO_PIN || info == base.INFO_CHECK {
				// We have detected a path that pins or checks. It will now
				// be marked in the info board, so move generation will skip
				// illegal moves.
				b.Squares[sliderSq.ToInfoIndex()] = info
				for stepSq := base.Square(int8(kingSq) + diffdir); stepSq != sliderSq; stepSq = base.Square(int8(stepSq) + diffdir) {
					b.Squares[stepSq.ToInfoIndex()] = info
				}
			}

			if checkCounter+ccount > 1 {
				// Double check detected. Leave early.
				return checkCounter
			}
		}
	}

	return checkCounter
}

// GeneratePawnMoves generates all pseudo legal pawn moves for the given color
// and stores them in the given MoveList.
func (b *Board) GeneratePawnMoves(mlist *base.MoveList, color base.Color) {
	from, to := base.OTB, base.OTB
	tpiece := base.NO_PIECE
	oppColor := color.Flip()
	move := base.Move{}
	piece := base.PAWN | color

	// Possible en passent captures are detected backwards
	// so we do not need to add another conditional to the
	// capture loop below.
	if b.EpSquare != base.OTB {
		from = b.EpSquare
		// We use the 'wrong' color to find e.p. captures
		// by searching in the opposite direction.
		// A potential overflow of int8(from) + capdir can generally be ignored
		// because casting it back to uint8 will correct the result.
		to = base.Square(int8(from) + base.PAWN_CAPTURE_DIRS[oppColor][0])
		tpiece = b.Squares[to]
		if to.OnBoard() && tpiece.HasColor(color) {
			move = base.NewMove(to, from, piece, base.PAWN|oppColor, base.NO_PIECE, base.EP_TYPE_CAPTURE, base.CASTLE_TYPE_NONE)
			mlist.Put(move)
		}

		to = base.Square(int8(from) + base.PAWN_CAPTURE_DIRS[oppColor][1])
		tpiece = b.Squares[to]
		if to.OnBoard() && tpiece.HasColor(color) {
			move = base.NewMove(to, from, piece, base.PAWN|oppColor, base.NO_PIECE, base.EP_TYPE_CAPTURE, base.CASTLE_TYPE_NONE)
			mlist.Put(move)
		}
	}

	for i := uint8(0); i < b.Pawns[color].Size; i++ {
		// 0. Retrieve 'from' square from piece list.
		from = b.Pawns[color].Pieces[i]

		// a. Captures
		for _, capdir := range base.PAWN_CAPTURE_DIRS[color] {
			to = base.Square(int8(from) + capdir)
			tpiece = b.Squares[to]
			// If the target square is on board and has the opponent's color
			// the capture is possible.
			if to.OnBoard() && tpiece.HasColor(oppColor) {
				// We also have to check if the capture is also a promotion.
				if to.IsPawnPromoting(color) {
					for prom := base.QUEEN; prom >= base.KNIGHT; prom >>= 1 {
						move = base.NewMove(from, to, piece, tpiece, prom|color, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
						mlist.Put(move)
					}
				} else {
					move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
					mlist.Put(move)
				}
			}
		}

		// b. Push by one.
		// Target square does never need legality check here.
		to = base.Square(int8(from) + base.PAWN_PUSH_DIRS[color])
		if b.Squares[to].IsEmpty() {
			if to.IsPawnPromoting(color) {
				for prom := base.QUEEN; prom >= base.KNIGHT; prom >>= 1 {
					move = base.NewMove(from, to, piece, base.NO_PIECE, prom|color, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
					mlist.Put(move)
				}
			} else {
				move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
				mlist.Put(move)
			}
		}
		// c. Double push by advancing one more time, if the pawn was at base rank.
		if from.IsPawnBaseRank(color) {
			to = base.Square(int8(to) + base.PAWN_PUSH_DIRS[color])
			if b.Squares[to].IsEmpty() {
				move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_CREATE, base.CASTLE_TYPE_NONE)
				mlist.Put(move)
			}
		}
	}
}

// GenerateKnightMoves generates all pseudo legal knight moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateKnightMoves(mlist *base.MoveList, color base.Color) {
	from, to := base.OTB, base.OTB
	tpiece := base.NO_PIECE
	oppColor := color.Flip()
	move := base.Move{}
	piece := base.KNIGHT | color

	// Iterate all knights of 'color'.
	for i := uint8(0); i < b.Knights[color].Size; i++ {
		from = b.Knights[color].Pieces[i]
		// Try all possible directions for a knight.
		for _, dir := range base.KNIGHT_DIRS {
			to = base.Square(int8(from) + dir)
			if to.OnBoard() {
				tpiece = b.Squares[to]
				if tpiece.IsEmpty() {
					// Add a normal move.
					move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
					mlist.Put(move)
				} else if tpiece.HasColor(oppColor) {
					// Add a capture move.
					move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
					mlist.Put(move)
				} // Else the square is occupied by own piece.
			} // Else target is outside of board.
		}
	}
}

// GenerateKingMoves generates all pseudo legal king moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateKingMoves(mlist *base.MoveList, color base.Color) {
	from, to := b.Kings[color], base.OTB
	tpiece := base.NO_PIECE
	oppColor := color.Flip()
	move := base.Move{}
	piece := base.KING | color

	if b.CheckInfo != CHECK_NONE {
		// Cannot castle when in check.
		goto SKIP_ALL_CASTLING
	}
	// Try castling first.
	// a. king-side
	// If castling king-side is still allowed..
	if b.CastleShort[color] {
		// test if the squares on short castling path are clear.
		for _, sq := range base.CASTLING_PATH_SHORT[color] {
			tpiece = b.Squares[sq]
			if !tpiece.IsEmpty() {
				goto SKIP_CASTLING_SHORT
			}
		}
		// All squares on the path are clear -> add castle short move.
		to = base.CASTLING_PATH_SHORT[color][1]
		move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_SHORT)
		mlist.Put(move)
	}

SKIP_CASTLING_SHORT:

	// b. queen-side
	// If castling queen side is still allowed..
	if b.CastleLong[color] {
		// test if the squares on long castling path are clear.
		for _, sq := range base.CASTLING_PATH_LONG[color] {
			tpiece = b.Squares[sq]
			if !tpiece.IsEmpty() {
				goto SKIP_CASTLING_LONG
			}
		}
		// All squares on the path are clear -> add castle short move.
		to = base.CASTLING_PATH_LONG[color][1]
		move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_LONG)
		mlist.Put(move)
	}

SKIP_CASTLING_LONG:
SKIP_ALL_CASTLING:

	// Now add all normal moves.
	for _, dir := range base.KING_DIRS {
		to = base.Square(int8(from) + dir)
		if to.OnBoard() {
			tpiece = b.Squares[to]
			if tpiece.IsEmpty() {
				// ADD NORMAL MOVE
				move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
				mlist.Put(move)
			} else if tpiece.HasColor(oppColor) {
				// ADD CAPTURE MOVE
				move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
				mlist.Put(move)
			}
		}
	}
}

// GenerateBishopMoves generates all pseudo legal bishop moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateBishopMoves(mlist *base.MoveList, color base.Color) {
	b.GenerateSlidingMoves(mlist, color, base.BISHOP, base.DIAGONAL_DIRS, &b.Bishops[color])
}

// GenerateRookMoves generates all pseudo legal rook moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateRookMoves(mlist *base.MoveList, color base.Color) {
	b.GenerateSlidingMoves(mlist, color, base.ROOK, base.ORTHOGONAL_DIRS, &b.Rooks[color])
}

// GenerateQueenMoves generates all pseudo legal rook moves for the given color
// and stores them in the given MoveList.
func (b *Board) GenerateQueenMoves(mlist *base.MoveList, color base.Color) {
	b.GenerateSlidingMoves(mlist, color, base.QUEEN, base.ORTHOGONAL_DIRS, &b.Queens[color])
	b.GenerateSlidingMoves(mlist, color, base.QUEEN, base.DIAGONAL_DIRS, &b.Queens[color])
}

// GenerateSlidingMoves generates all pseudo legal sliding moves for the given color
// and stores them in the given MoveList. This can be diagonal or orthogonal moves.
// This function is used to create all bishop, rook and queen moves.
func (b *Board) GenerateSlidingMoves(mlist *base.MoveList, color base.Color, ptype base.Piece, dirs [4]int8, plist *base.PieceList) {
	from, to := base.OTB, base.OTB
	tpiece := base.NO_PIECE
	move := base.Move{}
	piece := ptype | color

	// Iterate all specific pieces.
	for i := uint8(0); i < plist.Size; i++ {
		from = plist.Pieces[i]

		// For every direction a bishop can go..
		for _, dir := range dirs {
			// -> repeat until a stop condition occurs.

			for steps := int8(1); ; steps++ {
				to = base.Square(int8(from) + dir*steps)
				if !to.OnBoard() {
					// Target is an illegal square -> next direction.
					break
				} else {
					tpiece = b.Squares[to]
					if tpiece.HasColor(color) {
						// Target is blocked by own piece -> next direction.
						break
					} else {
						if tpiece.IsEmpty() {
							// Add a normal move.
							move = base.NewMove(from, to, piece, base.NO_PIECE, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
							mlist.Put(move)
							// And continue in current direction.
						} else {
							// Add a capture move.
							move = base.NewMove(from, to, piece, tpiece, base.NO_PIECE, base.EP_TYPE_NONE, base.CASTLE_TYPE_NONE)
							mlist.Put(move)
							// And go to next direction.
							break
						}
					}
				}

			}
		}

	}
}

func (b *Board) InfoBoardString() string {
	str := "  +-----------------+\n"
	for r := 7; r >= 0; r-- {
		str += strconv.Itoa(r+1) + " | "
		for f := 0; f < 8; f++ {
			idx := 16*r + f + 8
			str += fmt.Sprintf("%d ", b.Squares[idx])
			//			str += strconv.Itoa(int(b.Squares[idx])) + " "
			if f == 7 {
				str += "|\n"
			}
		}
	}
	str += "  +-----------------+\n"
	str += "    a b c d e f g h\n"

	return str
}

func (b *Board) String() string {
	str := "  +-----------------+\n"
	for r := 7; r >= 0; r-- {
		str += strconv.Itoa(r+1) + " | "
		for f := 0; f < 8; f++ {
			idx := 16*r + f
			if base.Piece(idx) == b.EpSquare {
				str += ", "
			} else {
				str += base.PrintMap[b.Squares[idx]] + " "
			}

			if f == 7 {
				str += "|\n"
			}
		}
	}
	str += "  +-----------------+\n"
	str += "    a b c d e f g h\n"

	return str
}
