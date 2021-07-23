package main

import (
	"fmt"
	"log"

	"github.com/clfs/mariposa/core"
)

func main() {
	b, err := core.NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Pretty())
	b, err = core.NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Pretty())
	b, err = core.NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Pretty())
	b, err = core.NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Pretty())
}
