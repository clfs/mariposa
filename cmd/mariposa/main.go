package main

import (
	"fmt"
	"log"

	"github.com/clfs/mariposa/core"
)

func main() {

	/*b, err := core.NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
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
	fmt.Println(b.Pretty())*/

	a := core.Board{Pieces: [64]core.Piece{0x5, 0xb, 0x6, 0xb, 0x4, 0x4, 0x4, 0x4, 0x5, 0x3, 0x0, 0x4, 0x5, 0x6, 0x1, 0xd, 0xb, 0xe, 0x2, 0x1, 0x6, 0x0, 0xc, 0xd, 0x6, 0x1, 0x5, 0xd, 0x0, 0x3, 0x0, 0x5, 0x6, 0x3, 0x3, 0x6, 0x3, 0x1, 0x0, 0xd, 0x3, 0x4, 0xb, 0xd, 0x2, 0x3, 0xa, 0x4, 0xd, 0x1, 0x1, 0xc, 0xb, 0x6, 0x6, 0x4, 0xc, 0x2, 0x5, 0x1, 0x6, 0x5, 0xc, 0x9}, SideToMove: 0x1, CastleRights: core.Castling{WhiteOO: true, WhiteOOO: false, BlackOO: false, BlackOOO: false}, EPTarget: 0x7d, HalfMoveClock: 0x922d34dc7e72dd04, FullMoveNumber: 0x24ac2ac375f9e051}
	b, err := core.NewBoard(a.FEN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", a)
	fmt.Println(a.FEN())
	fmt.Printf("%#v\n", b)
	fmt.Println(b.FEN())
}
