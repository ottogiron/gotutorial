package muni

import "fmt"

type Transport struct {
}

func (t *Transport) Ring() {
	fmt.Print("Ringing the bell... \n")
}
