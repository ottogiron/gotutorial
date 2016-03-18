package main

import (
	"github.com/ottogiron/gotutorial/muni"
	"github.com/ottogiron/gotutorial/workers"
)

var numbera = 1

func main() {
	bus := muni.NewBus("Azul")
	driver := workers.Driver{}
	driver.Drive(bus)
}
