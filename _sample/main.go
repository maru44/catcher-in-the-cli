package main

import (
	"fmt"
	"time"
)

func main() {
	// ctx := context.Background()
	ch := make(chan string)
	ms := []string{}

	// go catcher.Catch(ctx, ch, ms)

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
				if v != "" {
					ms = append(ms, v)
				}
			case <-time.After(3 * time.Second):
				// cancel()
				fmt.Println(ms)
				break
			}
		}
	}()

	fmt.Println("bbbb")
	fmt.Println("ccc")
}
