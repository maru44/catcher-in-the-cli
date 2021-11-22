package catcher

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"time"
)

func (c *catcher) CatchWithCtx(ctx context.Context, f func(ms []*Caught)) {
	localCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(c.Interval))
	defer cancel()

	chOut := make(chan bool)
	chIn := make(chan bool)
	chError := make(chan bool)

	if c.OutBulk != nil {
		go c.catchStdout(localCtx, chOut)
	}
	if c.InBulk != nil {
		go c.catchStdin(localCtx, chIn)
	}
	if c.ErrorBulk != nil {
		go c.catchStderr(localCtx, chError)
	}

	for {
		select {
		case <-localCtx.Done():
			for {
				if c.IsOver(chOut, chIn, chError) {
					cs := c.Separate()
					f(cs)
					c.Reset()
					return
				}
			}
		case <-ctx.Done():
			for {
				if c.IsOver(chOut, chIn, chError) {
					cs := c.Separate()
					f(cs)
					c.Reset()
					return
				}
			}
		}
	}
}

func (c *catcher) catchStdout(ctx context.Context, ch chan bool) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	for {
		select {
		case <-ctx.Done():
			w.Close()

			var buf bytes.Buffer
			io.Copy(&buf, r)

			c.OutBulk.Text = buf.String()

			os.Stdout = stdout // restore stdout
			ch <- true
			return
		}
	}
}

func (c *catcher) catchStderr(ctx context.Context, ch chan bool) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stderr := os.Stderr
	os.Stderr = w

	for {
		select {
		case <-ctx.Done():
			w.Close()

			var buf bytes.Buffer
			io.Copy(&buf, r)

			c.ErrorBulk.Text = buf.String()

			os.Stderr = stderr
			ch <- true
			return
		}
	}
}

func (c *catcher) catchStdin(ctx context.Context, ch chan bool) {
	c.InBulk.Text = ""
	s := bufio.NewScanner(os.Stdin)

	go func() {
		// for s.Scan() && !<-ch {
		// 	c.InBulk.Text += s.Text()
		// }

		for {
			select {
			case <-ctx.Done():
				ch <- true
				return
			}
		}
	}()

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		ch <- true
	// 		return
	// 	}
	// }

	for s.Scan() && !<-ch {
		c.InBulk.Text += s.Text()
	}
}

func (c *catcher) Separate() []*Caught {
	ret := []*Caught{}
	if c.OutBulk != nil {
		strs := strings.Split(c.OutBulk.Text, c.Separator)
		for _, s := range strs {
			if s != "" {
				ret = append(ret, &Caught{
					Text: s,
					Type: StdTypeOut,
				})
			}
		}
	}
	if c.InBulk != nil {
		strs := strings.Split(c.InBulk.Text, c.Separator)
		for _, s := range strs {
			if s != "" {
				ret = append(ret, &Caught{
					Text: s,
					Type: StdTypeIn,
				})
			}
		}
	}
	if c.ErrorBulk != nil {
		strs := strings.Split(c.ErrorBulk.Text, c.Separator)
		for _, s := range strs {
			if s != "" {
				ret = append(ret, &Caught{
					Text: s,
					Type: StdTypeError,
				})
			}
		}
	}
	return ret
}

// reset RawCaught
func (c *catcher) Reset() {
	if c.OutBulk != nil {
		c.OutBulk.Text = ""
	}
	if c.InBulk != nil {
		c.InBulk.Text = ""
	}
	if c.ErrorBulk != nil {
		c.ErrorBulk.Text = ""
	}
}

// is over all child
func (c *catcher) IsOver(chOut, chIn, chError chan bool) bool {
	if c.OutBulk != nil {
		if !<-chOut {
			return false
		}
	}
	if c.InBulk != nil {
		if !<-chIn {
			return false
		}
	}
	if c.ErrorBulk != nil {
		if !<-chError {
			return false
		}
	}
	return true
}
