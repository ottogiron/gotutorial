package muni

import "fmt"

type Metro struct {
	Transport
}

func (m *Metro) Move() {
	fmt.Print("Moving metro..\n")
}
