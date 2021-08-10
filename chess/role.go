package chess

import "fmt"

//go:generate stringer -type=Role
type Role uint8

const (
	Pawn Role = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

func (r Role) Value() uint8 {
	return uint8(r)
}

func (r Role) Valid() bool {
	return r >= Pawn && r <= King
}

func (r Role) Invalid() bool {
	return !r.Valid()
}

func ParseRole(s string) (Role, error) {
	switch s {
	case "P", "p":
		return Pawn, nil
	case "N", "n":
		return Knight, nil
	case "B", "b":
		return Bishop, nil
	case "R", "r":
		return Rook, nil
	case "Q", "q":
		return Queen, nil
	case "K", "k":
		return King, nil
	default:
		return 0, fmt.Errorf("todo invalid role error")
	}
}
