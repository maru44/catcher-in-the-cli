package catcher

import (
	"fmt"
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

func SetupCatcher(sec ...int) *Catcher {
	t := 60
	if len(sec) > 0 {
		t = sec[0]
	}
	return &Catcher{
		threadTime: t,
	}
}

func (m *MessageWithType) String() string {
	return fmt.Sprintf("%s: %v: %s", m.Type, m.Time, m.Message)
}
