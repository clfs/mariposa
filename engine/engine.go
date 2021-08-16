package engine

const (
	Name   = "Mariposa"
	Author = "Calvin Figuereo-Supraner <mail@calvin.page>"
)

type Engine struct {
}

func New() *Engine {
	return new(Engine)
}
