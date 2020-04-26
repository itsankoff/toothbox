// Package repl implements REPL (Read-Eval-Print loop)
package repl

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
)

var (
	// ErrQuit is returned if the .quit command is called
	ErrQuit = errors.New("repl: quit")
)

// Repl is implementation of REPL (Read-Eval-Print Loop)
type Repl struct {
	input  *bufio.Reader
	prefix string
	echo   bool
}

// New creates a new Repl instance
func New() *Repl {
	input := bufio.NewReader(os.Stdin)
	prefix := "toothbox>"

	return &Repl{
		input:  input,
		prefix: prefix + " ",
	}
}

// Run runs the repl
func (r *Repl) Run(ctx context.Context) error {
	for {
		select {
		case _ = <-ctx.Done():
			return nil
		default:
			fmt.Print(r.prefix)
			line, err := r.input.ReadString('\n')
			if err != nil {
				return err
			}

			if err := r.handleInput(line[:len(line)-1]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Repl) handleInput(in string) error {
	if r.echo {
		fmt.Println(in)
	}

	switch in {
	case "@echo on":
		r.echo = true
	case "@echo off":
		r.echo = false
	case ".help":
		r.printHelp()
	case ".quit":
		return ErrQuit
	default:
		fmt.Printf("Unrecognized command: %s\n", in)
	}

	return nil
}

func (r *Repl) printHelp() {
	fmt.Println(`
.help - print help
.quit - quit program
@echo on - turn on echo
@echo off - turn off echo
`)
}
