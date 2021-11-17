package catcher

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func InitCatcher(ctx context.Context) {
	ch := make(chan string)

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println("d")
				fmt.Println(v)
			}
		}
	}()
	catch(ctx, ch)
}

func (c *Catcher) catch(ctx context.Context, ch chan string) {
	localCtx, _ := context.WithCancel(ctx)
	c.initThreadTime()

	ms := []MessageWithType{}

	go scan(localCtx, os.Stdin, ms)
	go scan(localCtx, os.Stderr, ms)
	go scan(localCtx, os.Stdout, ms)
}

//
func scan(ctx context.Context, file *os.File, ms []MessageWithType) {
	scanner := bufio.NewScanner(file)

	var t StdType
	switch file {
	case os.Stdin:
		t = StdTypeIn
	case os.Stderr:
		t = StdTypeError
	default:
		t = StdTypeOut
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for scanner.Scan() {
				ms = append(ms, MessageWithType{
					Type:    t,
					Message: scanner.Text(),
					Time:    time.Now(),
				})
			}
		}
	}
}
