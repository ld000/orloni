package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ld000/orloni/player"
	"github.com/ld000/orloni/util"
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Main 程序执行
type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewMain return new Main
func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Run CLI 运行
func (m *Main) Run(args ...string) error {
	name, args := util.ParseCommandArgs(args)

	switch name {
	case "", "play":
		if err := player.Play(args); err != nil {
			return fmt.Errorf("run: %s", err)
		}
	default:
		return fmt.Errorf(`unknown command "%s"`+"\n\n", name)
	}
	return nil
}
