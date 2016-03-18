package muni

import "fmt"

//Bus a muni bus
type bus struct {
	Color string
	Transport
}

func NewBus(color string) Vehicle {
	return &bus{Color: color}
}

//Move the bus starts moving
func (b *bus) Move() {
	fmt.Print("Moving a bus ...\n")
}

func (b *bus) Ring() {
	fmt.Print("Ringing the bus bell....\n")
}
