// Package repl implements REPL (Read-Eval-Print loop)
package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// Repl is implementation of REPL (Read-Eval-Print Loop)
type Repl struct {
	input  *bufio.Reader
	prefix string
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

			fmt.Print(line)
		}
	}

	return nil
}
