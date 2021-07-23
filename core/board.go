package core

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Pieces         [NumSquares]Piece
	SideToMove     Color
	Castles        Castling
	EPTarget       Square
	HalfMoveClock  uint64
	FullMoveNumber uint64
}

// NewBoard returns a new board from a FEN.
func NewBoard(fen string) (Board, error) {
	board := Board{}

	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return Board{}, &InvalidFENError{fen}
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
			return Board{}, &InvalidFENError{fen}
		}
	}

	// 2. Side to move.
	switch fields[1] {
	case "w":
		board.SideToMove = White
	case "b":
		board.SideToMove = Black
	default:
		return Board{}, &InvalidFENError{fen}
	}

	// 3. Castling availability.
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			switch ru {
			case 'K':
				board.Castles.WhiteOO = true
			case 'Q':
				board.Castles.WhiteOOO = true
			case 'k':
				board.Castles.BlackOO = true
			case 'q':
				board.Castles.BlackOOO = true
			default:
				return Board{}, &InvalidFENError{fen}
			}
		}
	}

	// 4. En passant target square.
	board.EPTarget = SquareFromString(fields[3])

	// 5. Half move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return Board{}, &InvalidFENError{fen}
	}
	board.HalfMoveClock = uint64(halfMoveClock)

	// 6. Full move number.
	fullMoveNumber, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return Board{}, &InvalidFENError{fen}
	}
	board.FullMoveNumber = uint64(fullMoveNumber)

	return board, nil
}

func (b Board) PieceAt(s Square) Piece {
	return b.Pieces[s]
}

func (b Board) PieceAtFileRank(f File, r Rank) Piece {
	return b.PieceAt(SquareFromFileRank(f, r))
}

func (b Board) FEN() string {
	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		num := 0
		for f := FileA; f <= FileH; f++ {
			piece := b.PieceAtFileRank(f, r)
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
	sb.WriteString(b.SideToMove.String())
	sb.WriteString(" ")
	sb.WriteString(b.EPTarget.String())
	sb.WriteString(" ")
	sb.WriteString(b.Castles.String())
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(b.HalfMoveClock, 10))
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(b.FullMoveNumber, 10))
	sb.WriteString(" ")

	return sb.String()
}

func (b Board) Pretty() string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			fmt.Fprintf(&sb, "%s ", b.PieceAtFileRank(f, r))
		}
		sb.WriteString("\n")
	}

	fmt.Fprintf(&sb, "Side to move: %s\n", b.SideToMove.DebugString())
	fmt.Fprintf(&sb, "En passant target: %s\n", b.EPTarget.DebugString())
	fmt.Fprintf(&sb, "Half move clock: %d\n", b.HalfMoveClock)
	fmt.Fprintf(&sb, "Full move number: %d\n", b.FullMoveNumber)
	fmt.Fprintf(&sb, "Castles: %v\n", b.Castles)
	fmt.Fprintf(&sb, "FEN: %s", b.FEN())

	return sb.String()
}
