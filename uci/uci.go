// Package uci implements a UCI client specifically for the Mariposa engine.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/clfs/mariposa/chess"
	"github.com/clfs/mariposa/engine"
)

type Client struct {
	r io.Reader
	w io.Writer
	e *engine.Engine
	m sync.Mutex // protects w
}

func New(r io.Reader, w io.Writer) *Client {
	e := engine.New()
	return &Client{r: r, w: w, e: e}
}

func (c *Client) Run() error {
	s := bufio.NewScanner(c.r)
	for s.Scan() {
		if err := c.dispatch(s.Text()); err != nil {
			return err
		}
	}
	return s.Err()
}

func (c *Client) dispatch(command string) error {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return nil // ignore empty lines
	}
	first, rest := fields[0], fields[1:]

	switch first {
	case "uci":
		return c.uci()
	case "isready":
		return c.isReady()
	case "ucinewgame":
		return c.uciNewGame()
	case "position":
		return c.position(rest)
	case "go":
		return c.goCommand(rest)
	case "quit":
		return fmt.Errorf("todo: quit error message")
	default:
		return fmt.Errorf("unknown command error todo")
	}
}

func (c *Client) uci() error {
	log.Println("reached")
	c.m.Lock()
	fmt.Fprintf(c.w, "id name %s\n", engine.Name)
	fmt.Fprintf(c.w, "id author %s\n", engine.Author)
	fmt.Fprintf(c.w, "uciok\n")
	c.m.Unlock()
	return nil
}

func (c *Client) isReady() error {
	for !c.e.IsReady() {
		// wait. todo: better implementation than polling, since this will spam
	}
	c.m.Lock()
	fmt.Fprintf(c.w, "readyok\n")
	c.m.Unlock()
	return nil
}

func (c *Client) uciNewGame() error {
	c.e.NewGame()
	return nil
}

func (c *Client) position(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("todo, not enough args")
	}

	if args[0] == "fen" {
		fen := strings.Join(args[1:], "")
		return c.e.SetFEN(fen)
	}

	if args[0] == "startpos" {
		for _, move := range args[1:] {
			m, err := chess.ParseMove(move)
			if err != nil {
				return err
			}
			if err := c.e.Move(m); err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf("todo, error 182348172")
}

func (c *Client) goCommand(args []string) error {
	_ = args // todo implement arg handling
	bestMove, err := c.e.BestMove()
	if err != nil {
		return err
	}
	c.m.Lock()
	fmt.Fprintf(c.w, "bestmove %s\n", bestMove)
	c.m.Unlock()
	return nil
}
