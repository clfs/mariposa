package common

type EnPassantRight struct {
	target  Square
	allowed bool
}

func NewEnPassantRight() *EnPassantRight {
	return &EnPassantRight{}
}

func (e *EnPassantRight) FEN() string {
	if !e.allowed {
		return "-"
	}
	return e.target.FEN()
}

func (e *EnPassantRight) Target() (Square, bool) {
	return e.target, e.allowed
}

func (e *EnPassantRight) Set(target Square) {
	e.target = target
	e.allowed = true
}

func (e *EnPassantRight) Clear() {
	e.allowed = false
}
