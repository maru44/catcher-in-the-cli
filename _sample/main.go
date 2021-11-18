package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// ctx := context.Background()
	// ch := make(chan string)
	// // go catcher.InitCatcher(ctx)

	// c := catcher.SetupCatcher(5)

	// go c.Catch(ctx, ch)

	// time.Sleep(1 * time.Second)
	// fmt.Println("aaa")

	// time.Sleep(1 * time.Second)
	// fmt.Println("bbb")

	// time.Sleep(5 * time.Second)

	/* pattern 2 */

	t := []string{}
	ch := make(chan string)

	go func() {
		sc := bufio.NewScanner(os.Stdout)
		for sc.Scan() {
			// fmt.Println("d")
			// t = append(t, sc.Text())
			ch <- sc.Text()
		}
	}()

	go func() {
		select {
		case v := <-ch:
			t = append(t, v)
		}
	}()

	fmt.Println("a")

	time.Sleep(3 * time.Second)
	fmt.Println(t)
}
