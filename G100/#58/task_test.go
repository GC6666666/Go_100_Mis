package _58

import (
	"context"
	"io"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
)

//	func read(r io.Reader) (int, error) {
//		count := 0
//		for {
//			b := make([]byte, 1024)
//			_, err := r.Read(b)
//			if err != nil {
//				if err == io.EOF {
//					break
//				}
//				return 0, err
//			}
//			count += task(b)
//
//		}
//		return count, nil
//	}
func task(b []byte) int {

}

func read(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	n := 10

	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		ch <- b[:n]
	}
	ctx := context.WithoutCancel(context.Background())
	ctx.Done()
	close(ch)
	wg.Wait()
	return int(count), nil
	n = runtime.GOMAXPROCS(0)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// 执行一些任务来计算http响应

}
