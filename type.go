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
	threadTime int
}

func (c *Catcher) initThreadTime() {
	if c.threadTime == 0 {
		c.threadTime = 60
	}
}
