package main

import (
	"context"
	"os"
	"os/signal"
)

// main entry to applicationgo mod
func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	select {
	case <-ctx.Done():

		os.Exit(1)
	}
}
