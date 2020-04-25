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

	go func() {
		r := repl.New()
		r.Run(ctx)
	}()

	<-ctx.Done()
}
