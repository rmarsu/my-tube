package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/VandiKond/vanerrors"
)

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		msgs     = make([]string, 0, len(c.funcs))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, f := range c.funcs {
			if err := f(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return vanerrors.NewWrap("shutdown cancelled", ctx.Err(), vanerrors.EmptyHandler)
	}

	if len(msgs) > 0 {
		var errStr string
		for _, msg := range msgs {
			errStr += "\n" + msg
		}
		return vanerrors.NewBasic("shutdown finished with error(s)", errStr, vanerrors.EmptyHandler)
	}

	return nil
}

type Func func(ctx context.Context) error
