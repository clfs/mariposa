package chess

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
