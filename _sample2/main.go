package main

// @TODO without context

import (
	"fmt"
	"os"
	"time"

	"github.com/maru44/catcher-in-the-cli"
)

func main() {
	ch := make(chan string)
	c := catcher.GenerateCatcher(
		&catcher.Settings{
			Interval: 4000,
		},
	)

	go c.Catch(ch, println)

	time.Sleep(500 * time.Microsecond)
	fmt.Println("bbb")
	fmt.Println("ccc")

	fmt.Fprintln(os.Stderr, "ddddd")

	for {
		select {
		case v := <-ch:
			if v == catcher.SignalRepeat {
				go c.Catch(ch, println)
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
