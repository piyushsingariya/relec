package chanx

import (
	"context"
	"math"
	"sync/atomic"

	"github.com/piyushsingariya/relec/safego"
	"golang.org/x/sync/errgroup"
)

type Stream[T any] struct {
	started      atomic.Bool
	pool         []chan T
	roundrobin   atomic.Int64
	count        int
	chanConsumer func(ctx context.Context, channel chan T) error
}

func New[T any](count int, consumer func(ctx context.Context, channel chan T) error) *Stream[T] {
	return NewWithBuffer(count, 0, consumer)
}

func NewWithBuffer[T any](count, buffer int, consumer func(ctx context.Context, channel chan T) error) *Stream[T] {
	pool := []chan T{}
	for range count {
		var ch chan T
		if buffer > 0 {
			ch = make(chan T, buffer)
		} else {
			ch = make(chan T)
		}

		pool = append(pool, ch)
	}

	return &Stream[T]{
		pool:         pool,
		count:        count,
		chanConsumer: consumer,
	}
}

func (v *Stream[T]) StartSteaming(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	for _, ch := range v.pool {
		group.Go(func() error {
			return v.chanConsumer(ctx, ch)
		})
	}
	defer v.Close()

	v.started.Store(true)
	return group.Wait()
}

func (v *Stream[T]) Close() {
	for _, ch := range v.pool {
		safego.Close(ch)
	}
}

func (v *Stream[T]) Push(item T) bool {
	for !v.started.Load() {
	}

	in := v.roundrobin.Add(1) % int64(v.count)
	if in == math.MaxInt64 {
		v.roundrobin.Store(0)
	}

	return safego.Insert(v.pool[in], item)
}

func Insert[T any](ch chan<- T, value T) bool {
	safeInsert := false
	func() {
		ch <- value
		safeInsert = true
	}()

	return safeInsert
}
