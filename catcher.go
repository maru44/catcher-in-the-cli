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
	// catch(ctx, ch)
}

func (c *Catcher) Catch(ctx context.Context, ch chan string) {
	localCtx, cancel := context.WithCancel(ctx)

	ms := []MessageWithType{}

	time.Sleep(time.Duration(c.threadTime) * time.Second)
	scan(localCtx, os.Stdin, ms)
	scan(localCtx, os.Stderr, ms)
	scan(localCtx, os.Stdout, ms)
	cancel()

	for _, m := range ms {
		fmt.Println(m.String())
	}
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
