package engine

import (
	"fmt"
	"math/rand"

	"github.com/clfs/mariposa/chess"
)

const (
	Name   = "Mariposa"
	Author = "Calvin Figuereo-Supraner <mail@calvin.page>"
)

type Engine struct {
}

func New() *Engine {
	return new(Engine)
}

func (e *Engine) NewGame() {

}

func (e *Engine) SetFEN(s string) error {
	_ = s
	return fmt.Errorf("todo: need to implement")
}

func (e *Engine) IsReady() bool {
	return true
}

func (e *Engine) Move(m chess.Move) error {
	_ = m
	return fmt.Errorf("todo: need to implement")
}

func (e *Engine) Stop() error {
	return nil
}

func (e *Engine) BestMove() (chess.Move, error) {
	moves := e.LegalMoves()
	if len(moves) == 0 {
		return chess.Move{}, fmt.Errorf("no legal moves")
	}
	// choose randomly, todo eval
	best := moves[rand.Intn(len(moves))]
	return best, nil
}

func (e *Engine) LegalMoves() []chess.Move {
	// todo implement
	return []chess.Move{}
}
