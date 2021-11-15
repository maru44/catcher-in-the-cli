package catcher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
)

func InitCatcher() (context.Context, chan interface{}) {
	ctx := context.Background()
	ch := make(chan interface{})

	r, w, err := os.Pipe()
	old := os.Stdout

	if err != nil {
		panic(err)
	}
	os.Stdout = w

	// go catch(ctx, ch)
	go catch(ctx, r, ch)

	// w.Close()
	os.Stdout = old
	fmt.Println("origin: ", old)
	return ctx, ch
}

// catch cli
// func catch(ctx context.Context, ch chan interface{}) {
// 	for {
// 		select {
// 		case v := <-ch:
// 			fmt.Println(reflect.TypeOf(v))
// 		}
// 	}
// }

func catch(ctx context.Context, r *os.File, ch chan interface{}) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		panic(err)
	}
	ch <- buf.String()
}
