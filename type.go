package catcher

import (
	"time"
)

type (
	MessageWithType struct {
		Message string    `json:"message"`
		Type    StdType   `json:"stdType"`
		Time    time.Time `json:"time"`
	}

	StdType string
)

const (
	StdTypeError = "Error"
	StdTypeOut   = "Out"
	StdTypeIn    = "In"

	FinText = "Fin Catcher in the cli"
)

type Catcher struct {
	Interval int64
	Boundary string
}

type CatcherMaterial struct {
	Interval       int64 // millisecond
	BoundarySignal string
}

type Sample struct {
	Text string
}

func GenerateCatcher(m *CatcherMaterial) *Catcher {
	if m == nil {
		return &Catcher{
			Interval: 60000,
			Boundary: "\n",
		}
	}
	return &Catcher{
		Interval: m.Interval,
		Boundary: m.BoundarySignal,
	}
}
