package main

// @TODO without context

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/maru44/catcher-in-the-cli"
// )

// func main() {
// 	ctx := context.Background()

// 	c := catcher.GenerateCatcher(
// 		&catcher.Settings{
// 			Interval: 2000,
// 		},
// 	)

// 	go c.CatchWithCtx(ctx, println)

// 	time.Sleep(300 * time.Microsecond)
// 	fmt.Println("bbb")
// 	fmt.Println("ccc")

// 	fmt.Fprintln(os.Stderr, "ddddd")

// 	time.Sleep(3 * time.Second)
// }

// func println(ts []*catcher.Caught) {
// 	for _, t := range ts {
// 		fmt.Println(t)
// 	}
// }
