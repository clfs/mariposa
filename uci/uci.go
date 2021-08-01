// Package uci implements a subset of the UCI protocol, along with certain
// extensions.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/clfs/mariposa/engine"
)

func Run(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		command := s.Text()
		if err := Execute(w, command); err != nil {
			return err
		}
	}
	return s.Err()
}

func Execute(w io.Writer, command string) error {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return nil // ignore empty lines
	}
	first, rest := fields[0], fields[1:]

	// todo: use register/dispatch instead of switch/case
	var handler func(w io.Writer, args []string) error
	switch first {
	case "uci":
		handler = UCI
	case "isready":
		handler = IsReady
	case "setoption":
		handler = SetOption
	case "ucinewgame":
		handler = UCINewGame
	case "position":
		handler = Position
	case "go":
		handler = Go
	case "stop":
		handler = Stop
	case "quit", "exit", "kill":
		handler = Quit
	default:
		handler = Unknown
	}

	return handler(w, rest)
}

func UCI(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "id name %s\n", engine.Name)
	fmt.Fprintf(w, "id author %s\n", engine.Author)
	fmt.Fprintf(w, "uciok\n")
	return nil
}

func IsReady(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "readyok\n")
	return nil
}

func SetOption(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "not implemented yet\n")
	return nil
}

func UCINewGame(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "not implemented yet\n")
	return nil
}

func Position(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "not implemented yet\n")
	return nil
}

func Go(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "not implemented yet\n")
	return nil
}

func Stop(w io.Writer, args []string) error {
	_ = args
	fmt.Fprintf(w, "not implemented yet\n")
	return nil
}

func Quit(w io.Writer, args []string) error {
	_, _ = w, args
	return fmt.Errorf("quit")
}

func Unknown(w io.Writer, args []string) error {
	_, _ = w, args
	return fmt.Errorf("unknown command")
}
