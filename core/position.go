package core

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Position struct {
	Pieces         [NumSquares]Piece
	SideToMove     Color
	CastleRights   Castling
	EPTarget       Square
	HalfMoveClock  uint64
	FullMoveNumber uint64
}

// NewPosition returns a new position from a FEN.
func NewPosition(fen string) (Position, error) {
	board := Position{}

	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return Position{}, &InvalidFENError{fen}
	}

	// 1. Piece placement.
	square := A8
	for _, ru := range fields[0] {
		switch ru {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			board.Pieces[square] = PieceFromRune(ru)
			square++
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(ru - '0')
		case '/':
			square -= 16
		default:
			return Position{}, &InvalidFENError{fen}
		}
	}

	// 2. Side to move.
	switch fields[1] {
	case "w":
		board.SideToMove = White
	case "b":
		board.SideToMove = Black
	default:
		return Position{}, &InvalidFENError{fen}
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
				return Position{}, &InvalidFENError{fen}
			}
		}
	}

	// 4. En passant target square.
	board.EPTarget = SquareFromString(fields[3])

	// 5. Half move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return Position{}, &InvalidFENError{fen}
	}
	board.HalfMoveClock = uint64(halfMoveClock)

	// 6. Full move number.
	fullMoveNumber, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return Position{}, &InvalidFENError{fen}
	}
	board.FullMoveNumber = uint64(fullMoveNumber)

	return board, nil
}

func (p Position) PieceAt(s Square) Piece {
	return p.Pieces[s]
}

func (p Position) FEN() string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		num := 0
		for f := FileA; f <= FileH; f++ {
			piece := p.PieceAt(SquareAt(f, r))
			if piece.IsEmpty() {
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
		if r != Rank1 {
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

	return sb.String()
}

func (p Position) Pretty() string {
	var b strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			fmt.Fprintf(&b, "%s ", p.PieceAt(SquareAt(f, r)))
		}
		b.WriteString("\n")
	}
	fmt.Fprintf(&b, "FEN: %s", p.FEN())

	return b.String()
}
