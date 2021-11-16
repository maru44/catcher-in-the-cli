package catcher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
)

func InitCatcher(ctx context.Context) {
	ch := make(chan string)

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println("d")
				fmt.Println(v)
			}
		}
	}()

	catch(ctx, ch)
}

func catch(ctx context.Context, ch chan string) {
	old := os.Stdout
	fmt.Println("start")
	_, cancel := context.WithCancel(ctx)
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		fmt.Println(buf.String())
		if buf.String() != "" {
			fmt.Println("kita")
			ch <- buf.String()
			w.Close()
			os.Stdout = old
			cancel()
		}
	}()
}
