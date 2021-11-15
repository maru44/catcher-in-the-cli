package catcher

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
