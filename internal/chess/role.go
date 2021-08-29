package chess

type Role uint8

const (
	Pawn Role = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

// Valid returns whether the role is valid.
func (r Role) Valid() bool {
	return Pawn <= r && r <= King
}
