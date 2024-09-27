package relec

import (
	"context"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type CtxFunc = func(ctx context.Context) error

func Concurrent[T any](ctx context.Context, array []T, concurrency int, execute func(ctx context.Context, one T, executionNumber int) error) error {
	executor, ctx := errgroup.WithContext(ctx)
	executor.SetLimit(concurrency)

	for idx, one := range array {
		// schedule an execution
		// hold loop till a slot is available
		executor.Go(func() error {
			return execute(ctx, one, idx+1)
		})
	}

	// block the execution
	return executor.Wait()
}

func ConcurrentF(ctx context.Context, functions ...CtxFunc) error {
	executor, ctx := errgroup.WithContext(ctx)

	for _, one := range functions {
		// schedule an execution
		executor.Go(func() error {
			return one(ctx)
		})
	}

	// block the execution
	return executor.Wait()
}

func ConcurrentC[T any](ctx context.Context, yield <-chan T, concurrency int, execute func(ctx context.Context, one T, sequence int64) error) error {
	executor, ctx := errgroup.WithContext(ctx)
	executor.SetLimit(concurrency)

	// Channel to signal that all tasks have been scheduled
	done := make(chan struct{})

	go func() {
		defer close(done)
		counter := atomic.Int64{} // execution count to track the executions
		for {
			select {
			case <-ctx.Done():
				return
			case one, ok := <-yield:
				if !ok {
					return
				}
				sequence := counter.Add(1)
				// schedule an execution
				executor.Go(func() error {
					return execute(ctx, one, sequence)
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
