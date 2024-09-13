package lake

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/piyushsingariya/relec/safego"
	"golang.org/x/sync/errgroup"
)

type Channel[T any] struct {
	cancel context.CancelFunc
	ch     chan T
}

func (ch *Channel[T]) Insert(value T) {
	ch.ch <- value
}

func (ch *Channel[T]) close() {
	ch.cancel()
	safego.Close(ch.ch)
}

type Lake[T any] struct {
	ctx    context.Context
	cancel context.CancelFunc

	closedChannels atomic.Int64
	buffer         int
	execfunc       func(one T) (exit bool, err error)
	pool           []*Channel[T]
	controller     *errgroup.Group
}

func NewLake[T any](ctx context.Context, buffer int, exec func(one T) (exit bool, err error)) *Lake[T] {
	ctx, cancel := context.WithCancel(ctx)
	controller, _ := errgroup.WithContext(ctx)

	return &Lake[T]{
		ctx:            ctx,
		cancel:         cancel,
		buffer:         buffer,
		execfunc:       exec,
		controller:     controller,
		closedChannels: atomic.Int64{},
	}
}

func (l *Lake[T]) NewStream() *Channel[T] {
	ctx, cancel := context.WithCancel(l.ctx)
	ch := &Channel[T]{
		cancel: cancel,
	}
	if l.buffer > 0 {
		ch.ch = make(chan T, l.buffer)
	} else {
		ch.ch = make(chan T)
	}

	l.controller.Go(func() error {
		defer l.closedChannels.Add(1)

		for {
			select {
			case <-ctx.Done():
				return nil
			case one := <-ch.ch:
				exit, err := l.execfunc(one)
				if err != nil {
					return err
				}
				// if exit then exit
				if exit {
					return nil
				}
			}
		}
	})

	l.pool = append(l.pool, ch)
	return ch
}

func (l *Lake[T]) Wait(cancel func(chan T) error) error {
	defer func() {
		for _, ch := range l.pool {
			ch.close()
		}
	}()

	for _, ch := range l.pool {
		cancel(ch.ch)
	}

	return l.controller.Wait()
}

func (l *Lake[T]) Stats() string {
	return fmt.Sprintf("Completed: %d, Total: %d", int(l.closedChannels.Load()), len(l.pool))
}
