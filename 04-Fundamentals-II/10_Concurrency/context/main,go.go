package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	doSimple(ctx, "Test")

	vctx := context.WithValue(ctx, "ID", "1234")
	doWithValue(vctx, "More testing")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	doWithTimeout(ctx, cancel)

}

func doSimple(ctx context.Context, s string) {
	fmt.Println(s)
}

func doWithValue(ctx context.Context, s string) {
	fmt.Println(ctx.Value("ID"), s)
}


func doWithTimeout(ctx context.Context, cancel context.CancelFunc) {

	go func(ctx context.Context) {
		time.Sleep(100 * time.Second)
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("timeout exceeded")
		case context.Canceled:
			fmt.Println("context cancelled")
		}
	}
}
