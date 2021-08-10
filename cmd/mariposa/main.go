package main

import (
	"log"
	"os"

	"github.com/clfs/mariposa/uci"
)

func main() {
	log.Fatal(uci.New(os.Stdin, os.Stdout).Run())
}
