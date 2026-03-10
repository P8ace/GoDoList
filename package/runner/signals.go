package runner

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalError struct {
	sig os.Signal
}

// Error implements the error interface.
func (e SignalError) Error() string {
	return fmt.Sprintf("received signal %s", e.sig)
}

// function to handle graceful shutdowns
func ListenInterrupts(ctx context.Context) (execute func() error, interrupt func(error)) {
	ctx, cancel := context.WithCancel(ctx)
	return func() error {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			defer signal.Stop(sigChan)
			select {
			case sig := <-sigChan:
				return SignalError{sig: sig}
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(err error) {
			cancel()
		}
}
