# The Catcher in the Cli

> *I'm standing on the edge of some crazy cliff. What I have to do, I have to catch everybody if they start to go over the cliff—I mean if they’re running and they don’t look where they’re going I have to come out from somewhere and catch them. That’s all I’d do all day. I’d just be the catcher in the cli and all.*

<br />
ref: *The Catcher in the Rye*

## Explain

You can catch `Stdout`, `Stdin` and `Stderr` by using this package.

## Usage

I'll show you usage and that result.

### usage

```go:main.go
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

	c := catcher.GenerateCatcher(
		&catcher.Settings{
			Interval: 4000,
			Repeat:   catcher.IntPtr(2),
		},
	)

	ch := make(chan string)

	go c.CatchWithCtx(ctx, ch, writeFile)

	// you need time blank
	time.Sleep(500 * time.Microsecond)

	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Fprintln(os.Stderr, "ddddd")

	for {
		select {
		case v := <-ch:
			if v == catcher.SignalRepeat {
				go c.CatchWithCtx(ctx, ch, writeFile)
			} else {
				return
			}
		}
	}
}

func writeFile(ts []*catcher.Caught) {
	f, _ := os.OpenFile("./_sample/log.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer f.Close()

	for _, t := range ts {
		f.Write([]byte(t.String() + "\n"))
	}
}

```

### result

I ran `go run _sample/main.go` from `<my_directory>/catcher-in-the-cli`.
<br />And also I ran following commands `aaa` and `ls` under main func is running.

**output**
```
aaa
ddddd
exec: "aaa": executable file not found in $PATH
bbb
ccc
ls
LICENSE
README.md
_sample
_sample2
catcher.go
caught.go
domain.go
go.mod
tools.go

```

**log.log**
```log.log
Output: bbb
Output: ccc
Input: aaa
Error: ddddd
Error: exec: "aaa": executable file not found in $PATH
Output: LICENSE
Output: README.md
Output: _sample
Output: _sample2
Output: catcher.go
Output: caught.go
Output: domain.go
Output: go.mod
Output: tools.go
Input: ls

```
