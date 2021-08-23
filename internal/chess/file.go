package chess

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

func (f *File) Mirror() {
	*f = 7 - *f
}
