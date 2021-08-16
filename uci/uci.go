// Package uci implements a UCI client specifically for the Mariposa engine.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

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
	first, _ := fields[0], fields[1:]

	switch first {
	case "uci":
		return c.uci()
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
