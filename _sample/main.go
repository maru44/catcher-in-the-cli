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
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ms := &catcher.Sample{
		Text: "a",
	}

	//

	go catcher.Catch(ctx, ms)

	fmt.Println("bbbb")
	fmt.Println("ccc")
	fmt.Fprintln(os.Stderr, "ddddd")

	time.Sleep(2 * time.Second)
	fmt.Println(ms)
}
