package hkgo

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestAsync(t *testing.T) {
	signal := make(chan struct{})
	Async(func() {
		signal <- struct{}{}
	}, func(ctx context.Context) error {
		fmt.Println("123")
		return nil
	}, func(ctx context.Context) error {
		fmt.Println("1234")
		return nil
	}, func(ctx context.Context) error {
		fmt.Println("1235")
		return nil
	}, func(ctx context.Context) error {
		fmt.Println("1236")
		return errors.New("1236")
	})
	_ = <-signal
}
