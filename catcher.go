package catcher

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

// func InitCatcher(ctx context.Context) {
// 	ch := make(chan string)

// 	go func() {
// 		for {
// 			select {
// 			case v := <-ch:
// 				fmt.Println("d")
// 				fmt.Println(v)
// 			}
// 		}
// 	}()
// 	catch(ctx, ch)
// }

func (c *Catcher) Catch(ctx context.Context, ch chan string, ms []string) {
	// ctx, cancel := context.WithCancel(ctx)
	// ms := []MessageWithType{}

	go func() {
		for {
			select {
			case v := <-ch:
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

	for {
		scan(ctx, ch, os.Stdin)
		scan(ctx, ch, os.Stdout)
		scan(ctx, ch, os.Stderr)
	}
}

//
func scan(ctx context.Context, ch chan string, file *os.File) {
	fmt.Println("a")
	// buf := bytes.Buffer{}
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// var t StdType
	// switch file {
	// case os.Stdin:
	// 	t = StdTypeIn
	// case os.Stderr:
	// 	t = StdTypeError
	// default:
	// 	t = StdTypeOut
	// }

	go func() {
		var b bytes.Buffer
		io.Copy(&b, r)
		ch <- b.String()
	}()

	w.Close()
}
