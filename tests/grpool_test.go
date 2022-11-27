package tests

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtimer"
	"testing"
	"time"
)

func job(ctx context.Context) {
	time.Sleep(1 * time.Second)
}

func TestPool(t *testing.T) {
	pool := grpool.New(100)
	for i := 0; i < 1000; i++ {
		pool.Add(ctx, job)
	}
	fmt.Println("worker:", pool.Size())
	fmt.Println(" jobs:", pool.Jobs())
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		fmt.Println("worker:", pool.Size())
		fmt.Println(" jobs:", pool.Jobs())
		fmt.Println()
	})

	select {}
}
