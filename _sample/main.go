package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/maru44/catcher-in-the-cli"
)

func main() {
	ctx := context.Background()

	c := catcher.GenerateCatcher(
		&catcher.Settings{
			Interval: 4000,
			Repeat:   catcher.IntPtr(1),
		},
	)

	ch := make(chan string)

	go c.CatchWithCtx(ctx, ch, println)

	time.Sleep(300 * time.Microsecond)
	fmt.Println("bbb")
	fmt.Println("ccc")

	fmt.Fprintln(os.Stderr, "ddddd")

	for {
		select {
		case v := <-ch:
			if v == catcher.SignalRepeat {
				go c.CatchWithCtx(ctx, ch, println)
			} else {
				return
			}
		}
	}
}

func println(ts []*catcher.Caught) {
	for _, t := range ts {
		fmt.Println(t)
	}
}
