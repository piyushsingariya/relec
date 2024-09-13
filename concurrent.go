package relec

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Concurrent[T any](ctx context.Context, array []T, concurrency int, execute func(one T) error) error {
	executor, ctx := errgroup.WithContext(ctx)
	exit, cancel := context.WithCancel(ctx)
	defer cancel()

	semaphore := make(chan int, concurrency)

	for _, one := range array {
		// schedule an execution
		executor.Go(func() error {
			// hold loop till a slot is available
		hold:
			for {
				select {
				case semaphore <- 1:
					break hold
				case <-exit.Done():
					return nil
				}
			}

			// schedule defering to enable an execution slot
			defer func() {
				<-semaphore
			}()
			return execute(one)
		})
	}

	// block the execution
	return executor.Wait()
}
