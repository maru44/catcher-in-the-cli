package main

import (
	"fmt"
	"time"

	"github.com/maru44/catcher-in-the-cli"
)

func main() {
	ctx, _ := catcher.InitCatcher()

	fmt.Println("aaa")

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("bbb")

	time.Sleep(5000 * time.Millisecond)
	ctx.Done()
}
