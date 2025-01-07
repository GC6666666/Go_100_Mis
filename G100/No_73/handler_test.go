package No_73

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {

}

func handler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		i := i
		circle := circle
		go func() {
			defer wg.Done()
			result, err := foo(ctx, circle)
			if err != nil {
				//?
			}
			results[i] = result
		}()
	}
	wg.Wait()
	return results, nil
}

func handler1(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

type Circle struct {
	len int
}

type Result struct {
	result int
}

func foo(ctx context.Context, circle Circle) (Result, error) {
}
