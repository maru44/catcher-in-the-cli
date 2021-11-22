package catcher

import (
	"context"
)

type (
	StdType string

	RawCaught struct {
		Text string  `json:"text"`
		Type StdType `json:"stdType"`
	}

	Caught struct {
		Text string  `json:"text"`
		Type StdType `json:"stdType"`
	}

	catcher struct {
		Settings
		OutBulk   *RawCaught
		InBulk    *RawCaught
		ErrorBulk *RawCaught
	}

	Settings struct {
		// millisecond
		Interval   int64
		Separator  string
		TargetType []StdType
	}

	Sample struct {
		Text string
	}

	CatcherInTheCli interface {
		// Setting(s *Settings) error
		Catch(f func(m *[]CaughtInTheCli))
		CatchWithCtx(ctx context.Context, f func(m *[]CaughtInTheCli))
		Separate() []*CaughtInTheCli
		Reset()
		IsOver(chOut, chIn, chError chan bool) bool
	}

	CaughtInTheCli interface {
		String() string
		Json() ([]byte, error)
	}
)

const (
	StdTypeError = "Error"
	StdTypeOut   = "Output"
	StdTypeIn    = "Input"
)

// generator of catcher
func GenerateCatcher(s *Settings) catcher {
	if s == nil {
		return catcher{
			Settings: Settings{
				Interval:  60000,
				Separator: "\n",
				TargetType: []StdType{
					StdTypeOut, StdTypeIn, StdTypeError,
				},
			},
			OutBulk:   &RawCaught{Type: StdTypeOut},
			InBulk:    &RawCaught{Type: StdTypeIn},
			ErrorBulk: &RawCaught{Type: StdTypeError},
		}
	}

	c := &catcher{}
	if s.TargetType == nil || len(s.TargetType) == 0 {
		s.TargetType = []StdType{StdTypeError, StdTypeIn, StdTypeOut}
		c.OutBulk = &RawCaught{Type: StdTypeOut}
		c.InBulk = &RawCaught{Type: StdTypeIn}
		c.ErrorBulk = &RawCaught{Type: StdTypeError}
	} else {
		for _, t := range s.TargetType {
			if t == StdTypeOut {
				c.OutBulk = &RawCaught{Type: StdTypeOut}
			}
			if t == StdTypeIn {
				c.InBulk = &RawCaught{Type: StdTypeIn}
			}
			if t == StdTypeError {
				c.ErrorBulk = &RawCaught{Type: StdTypeError}
			}
		}
	}

	// default interval
	if s.Interval == 0 {
		s.Interval = 60000
	}

	// default separator
	if s.Separator == "" {
		s.Separator = "\n"
	}

	c.Settings = *s
	return *c
}
