package main

import (
	"fmt"

	"github.com/ottogiron/gotutorial/muni"
)

var numbera = 1

func main() {
	bus := muni.Bus{
		Plate: "XRE-71",
		Color: "Azul",
		Brand: "Mercedez Benz",
	}
	var err error
	if err = bus.Move(); err != nil {
		fmt.Print(err)
	}

}
