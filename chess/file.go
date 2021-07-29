package chess

//go:generate stringer -type File
type File uint8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) Value() uint8 {
	return uint8(f)
}

func (f File) Valid() bool {
	return f <= FileH
}

func (f File) Invalid() bool {
	return !f.Valid()
}
