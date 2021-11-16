package catcher

import "os"

type (
	MessageWithType struct {
		Message string  `json:"message"`
		Type    StdType `json:"stdType"`
	}

	StdType string
)

const (
	StdTypeError = "Error"
	StdTypeOut   = "Out"
)

type Catcher struct {
	saved  *os.File
	writer *os.File
	reader *os.File
}
