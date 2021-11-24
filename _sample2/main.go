package main

// @TODO without context

import (
	"fmt"
	"os"
	"time"

	"github.com/maru44/catcher-in-the-cli"
)

func main() {
	c := catcher.GenerateCatcher(
		&catcher.Settings{
			Interval: 4000,
		},
	)

	go func() {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("bbb")
			fmt.Println("ccc")

			fmt.Fprintln(os.Stderr, "ddddd")
		}
	}()

	c.Catch(println)
}

func println(ts []*catcher.Caught) {
	for _, t := range ts {
		fmt.Println(t)
	}
}
