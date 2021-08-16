package common

type EnPassantRight struct {
	Target  Square
	Allowed bool
}

func NewEnPassantRight() *EnPassantRight {
	return &EnPassantRight{}
}

func (e *EnPassantRight) FEN() string {
	if !e.Allowed {
		return "-"
	}
	return e.Target.FEN()
}

func (e *EnPassantRight) Get() (Square, bool) {
	return e.Target, e.Allowed
}

func (e *EnPassantRight) Set(target Square) {
	e.Target = target
	e.Allowed = true
}

func (e *EnPassantRight) Clear() {
	e.Allowed = false
}

func (e *EnPassantRight) Equal(x EnPassantRight) bool {
	return (!e.Allowed && !x.Allowed) || (e.Target == x.Target)
}
