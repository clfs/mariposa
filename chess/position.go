package chess

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	Board Board
	State State
}

func NewPosition(fen string) (*Position, error) {
	p := new(Position)
	if err := p.setFEN(fen); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Position) setFEN(fen string) error {
	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return fmt.Errorf("todo parsefenerror")
	}

	// 1. Piece placement.
	square := A8
	for _, ru := range fields[0] {
		switch ru {
		case 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k':
			piece, err := ParsePiece(string(ru))
			if err != nil {
				return fmt.Errorf("todo %v", err)
			}
			p.Board.Put(piece, square)
			square++
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(ru - '0') // todo weird typing
		case '/':
			square -= 16
		default:
			return fmt.Errorf("todo bad rune %v", ru)
		}
	}

	// 2. Side to move.
	// todo maybe parse color func
	switch fields[1] {
	case "w":
		p.State.SideToMove = White
	case "b":
		p.State.SideToMove = Black
	default:
		return fmt.Errorf("todo bad side to move %v", fields[1])
	}

	// 3. Castling availability.
	p.State.Castling = 0 // this is required since position might have stale info
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			switch ru {
			case 'K':
				p.State.Castling.Set(WhiteOO)
			case 'Q':
				p.State.Castling.Set(WhiteOOO)
			case 'k':
				p.State.Castling.Set(BlackOO)
			case 'q':
				p.State.Castling.Set(BlackOOO)
			default:
				return fmt.Errorf("todo bad castling %v", ru)
			}
		}
	}

	// 4. En passant target square.
	if fields[3] != "-" {
		square, err := ParseSquare(fields[3])
		if err != nil {
			return fmt.Errorf("todo %v", err)
		}
		p.State.EnPassantTarget = square
	} else {
		p.State.EnPassantAllowed = false
	}

	// 5. Half-move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return fmt.Errorf("todo %v", err)
	}
	p.State.HalfMoveClock = halfMoveClock

	// 6. Full-move number.
	fullMoveNumber, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return fmt.Errorf("todo %v", err)
	}
	p.State.FullMoveCount = fullMoveNumber

	return nil
}

func (p *Position) FEN() (string, error) {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		num := 0
		for f := FileA; f <= FileH; f++ {
			square, err := SquareAt(f, r)
			if err != nil {
				return "", err
			}
			piece, ok := p.Board.Get(square)
			if !ok {
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
	sb.WriteString(p.State.SideToMove.String())
	sb.WriteString(" ")
	sb.WriteString(p.State.Castling.String())
	sb.WriteString(" ")
	if p.State.EnPassantAllowed {
		sb.WriteString(p.State.EnPassantTarget.String())
	} else {
		sb.WriteString("-")
	}
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(p.State.HalfMoveClock, 10))
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(p.State.FullMoveCount, 10))
	return sb.String(), nil
}

func (p *Position) Pretty() (string, error) {
	var b strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			square, err := SquareAt(f, r)
			if err != nil {
				return "", err
			}
			piece, ok := p.Board.Get(square)
			if !ok {
				b.WriteString(". ")
			} else {
				fmt.Fprintf(&b, "%s ", piece)
			}
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

func (p *Position) MaybeLegalMoves() []Move {
	return nil // todo
}

func (p *Position) LegalMoves() []Move {
	return nil // todo
}
