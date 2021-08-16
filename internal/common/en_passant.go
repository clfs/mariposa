package common

type EnPassantRight struct {
	target  Square
	allowed bool
}

// todo: move this to fen package
func ParseEnPassantRight(s string) (EnPassantRight, error) {
	if s == "-" {
		return EnPassantRight{}, nil
	}
	square, err := ParseSquare(s)
	if err != nil {
		return EnPassantRight{}, err
	}
	return EnPassantRight{target: square}, nil
}

func (e *EnPassantRight) FEN() string {
	if !e.allowed {
		return "-"
	}
	return e.target.FEN()
}

func (e *EnPassantRight) Allowed() bool {
	return e.allowed
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
