package relec

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Concurrent[T any](ctx context.Context, array []T, concurrency int, execute func(one T) error) error {
	executor, _ := errgroup.WithContext(ctx)
	executor.SetLimit(concurrency)

	for _, one := range array {
		// schedule an execution
		// hold loop till a slot is available
		executor.Go(func() error {
			return execute(one)
		})
	}

	// block the execution
	return executor.Wait()
}

func ConcurrentC[T any](ctx context.Context, yield <-chan T, concurrency int, execute func(one T) error) error {
	executor, ctx := errgroup.WithContext(ctx)
	executor.SetLimit(concurrency)

	// Channel to signal that all tasks have been scheduled
	done := make(chan struct{})

	go func(){
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			case one := <- yield:
				// schedule an execution
				executor.Go(func() error {
					return execute(one)
				})
			}
		}
	}()

	// block the execution
	select {
	case <-done:
		return executor.Wait()
	case <-ctx.Done():
		return executor.Wait()
	}
}

func Yield[T any](generate func(channel chan<- T)) <-chan T {
	channel := make(chan T)

	go func() {
		defer close(channel)
		generate(channel)
	}()

	return channel
}
