package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/itsankoff/toothbox/repl"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Printf("Signal received: %v\n", sig)
		cancel()
	}()

	var err error
	go func() {
		r := repl.New()
		err = r.Run(ctx)
		if err != nil {
			if err != repl.ErrQuit {
				fmt.Println(err)
			} else {
				err = nil
			}

			cancel()
		}
	}()

	<-ctx.Done()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
