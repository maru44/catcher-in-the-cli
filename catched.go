package catcher

import (
	"encoding/json"
	"fmt"
)

// method implemention of catched

func (m *Caught) String() string {
	return fmt.Sprintf("%s: %s", m.Type, m.Text)
}

func (m *Caught) Json() ([]byte, error) {
	return json.Marshal(m)
}
