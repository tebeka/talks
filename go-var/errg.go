package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, timeout time.Duration, fail bool) error {
	fmt.Println(id, "start")
	time.Sleep(timeout)
	fmt.Println(id, "done")
	if fail {
		fmt.Println(id, "fail")
		return fmt.Errorf("%d failed", id)
	}
	return nil
}

func errExample() {
	start := time.Now()

	// START OMIT
	var g errgroup.Group
	g.Go(func() error {
		return worker(1, time.Second, false)
	})
	// END OMIT

	g.Go(func() error {
		return worker(2, 2*time.Second, true)
	})
	g.Go(func() error {
		return worker(3, 3*time.Second, false)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(time.Since(start))
}

func ctxWorker(ctx context.Context, cancel func(), id int, timeout time.Duration, fail bool) error {
	fmt.Println(id, "START")
	select {
	case <-ctx.Done():
		fmt.Println(id, "TIMEOUT")
		return nil
	case <-time.After(timeout):
		fmt.Println(id, "DONE")
		if fail {
			fmt.Println(id, "FAIL")
			cancel()
			return fmt.Errorf("%d failed", id)
		}
		return nil
	}
}

func ctxExample() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	start := time.Now()
	g.Go(func() error {
		return ctxWorker(ctx, cancel, 1, time.Second, false)
	})
	g.Go(func() error {
		return ctxWorker(ctx, cancel, 2, 2*time.Second, true)
	})
	g.Go(func() error {
		return ctxWorker(ctx, cancel, 3, 3*time.Second, false)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(time.Since(start))
}

func main() {
	errExample()
	ctxExample()
}
