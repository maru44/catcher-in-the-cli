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
	ch := make(chan string)

	c := catcher.GenerateCatcher(
		&catcher.Settings{
			Interval: 4000,
			Repeat:   catcher.IntPtr(2),
		},
	)

	go c.CatchWithCtx(ctx, ch, writeFile)

	// you need time blank
	time.Sleep(500 * time.Microsecond)

	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Fprintln(os.Stderr, "ddddd")

	for {
		select {
		case v := <-ch:
			if v == catcher.SignalRepeat {
				go c.CatchWithCtx(ctx, ch, writeFile)
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

func writeFile(ts []*catcher.Caught) {
	f, _ := os.OpenFile("./_sample/log.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer f.Close()

	for _, t := range ts {
		f.Write([]byte(t.String() + "\n"))
	}
}
