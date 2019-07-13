package basic

import (
	"context"
	"log"
	"testing"
	"time"
)

// https://draveness.me/golang/concurrency/golang-context.html
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(duration):
		log.Println("process request with", duration)
	}
}

func TestContext1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		log.Println("TestContext", ctx.Err())
	}
}

func TestContext2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		log.Println("TestContext", ctx.Err())
	}
}
