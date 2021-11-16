package main

import (
	"context"
	"fmt"
	"time"

	"github.com/maru44/catcher-in-the-cli"
)

func main() {
	ctx := context.Background()
	go catcher.InitCatcher(ctx)

	time.Sleep(1 * time.Second)
	fmt.Println("aaa")

	time.Sleep(1 * time.Second)
	fmt.Println("bbb")
}
