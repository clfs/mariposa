package chess

import (
	"fmt"
	"strings"
)

type Position struct {
	board Board
	state State
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
			p.board.Put(piece, square)
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
		p.state.SideToMove = White
	case "b":
		p.state.SideToMove = Black
	default:
		return fmt.Errorf("todo bad side to move %v", fields[1])
	}

	// 3. Castling availability.
	p.state.Castling = 0 // this is required since position might have stale info
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			switch ru {
			case 'K':
				p.state.Castling.Set(WhiteOO)
			case 'Q':
				p.state.Castling.Set(WhiteOOO)
			case 'k':
				p.state.Castling.Set(BlackOO)
			case 'q':
				p.state.Castling.Set(BlackOOO)
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
		p.state.EnPassantTarget = square
	} else {
		p.state.EnPassantAllowed = false
	}

	// 5. Half-move clock.
	// todo

	// 6. Full-move number.
	// todo

	return nil
}

func (p *Position) Board() Board {
	return p.board
}
