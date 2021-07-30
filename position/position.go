package position

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/clfs/mariposa/chess"
)

type Position struct {
	Pieces         [64]chess.Piece
	SideToMove     chess.Color
	CastleRights   Castling
	EPTarget       chess.Square
	HalfMoveClock  uint64
	FullMoveNumber uint64
}

// NewPosition returns a new position from a FEN.
func NewPosition(fen string) (Position, error) {
	board := Position{}

	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return Position{}, &ParseFENError{fen}
	}

	// 1. Piece placement.
	square := chess.A8
	for _, ru := range fields[0] {
		switch ru {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			piece, err := chess.ParsePiece(string(ru))
			if err != nil {
				return Position{}, err
			}
			board.Pieces[square] = piece
			square++
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += chess.Square(ru - '0')
		case '/':
			square -= 16
		default:
			return Position{}, &ParseFENError{fen}
		}
	}

	// 2. Side to move.
	switch fields[1] {
	case "w":
		board.SideToMove = chess.White
	case "b":
		board.SideToMove = chess.Black
	default:
		return Position{}, &ParseFENError{fen}
	}

	// 3. Castling availability.
	board.CastleRights = Castling{}
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			log.Println(ru)
			switch ru {
			case 'K':
				board.CastleRights.WhiteOO = true
			case 'Q':
				board.CastleRights.WhiteOOO = true
			case 'k':
				board.CastleRights.BlackOO = true
			case 'q':
				board.CastleRights.BlackOOO = true
			default:
				return Position{}, &ParseFENError{fen}
			}
		}
	}

	// 4. En passant target square.
	square, err := chess.ParseSquare(fields[3])
	if err != nil {
		board.EPTarget = chess.Square(0xff) // TODO: fix
	} else {
		board.EPTarget = square
	}

	// 5. Half move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return Position{}, &ParseFENError{fen}
	}
	board.HalfMoveClock = uint64(halfMoveClock)

	// 6. Full move number.
	fullMoveNumber, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return Position{}, &ParseFENError{fen}
	}
	board.FullMoveNumber = uint64(fullMoveNumber)

	return board, nil
}

func (p Position) PieceAt(s chess.Square) chess.Piece {
	return p.Pieces[s]
}

func (p Position) FEN() (string, error) {
	var sb strings.Builder

	for r := chess.Rank8; r <= chess.Rank8; r-- {
		num := 0
		for f := chess.FileA; f <= chess.FileH; f++ {
			square, err := chess.SquareAt(f, r)
			if err != nil {
				return "", err
			}
			piece := p.PieceAt(square)
			if piece.Invalid() {
				num += 1
			} else {
				if num != 0 {
					sb.WriteString(strconv.Itoa(num))
				}
				sb.WriteString(piece.String())
				num = 0
			}
		}
		if num != 0 {
			sb.WriteString(strconv.Itoa(num))
		}
		if r != chess.Rank1 {
			sb.WriteString("/")
		}
	}

	sb.WriteString(" ")
	sb.WriteString(p.SideToMove.String())
	sb.WriteString(" ")
	sb.WriteString(p.CastleRights.String())
	sb.WriteString(" ")
	sb.WriteString(p.EPTarget.String())
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(p.HalfMoveClock, 10))
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(p.FullMoveNumber, 10))

	return sb.String(), nil
}

func (p Position) Pretty() (string, error) {
	var b strings.Builder

	for r := chess.Rank8; r <= chess.Rank8; r-- {
		for f := chess.FileA; f <= chess.FileH; f++ {
			square, err := chess.SquareAt(f, r)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&b, "%s ", p.PieceAt(square))
		}
		b.WriteString("\n")
	}
	fen, err := p.FEN()
	if err != nil {
		return "", err
	}
	fmt.Fprintf(&b, "FEN: %s", fen)

	return b.String(), nil
}

func (p Position) PseudoLegalMoves() []Move {
	return nil
}

func (p Position) LegalMoves() []Move {
	return nil
}
