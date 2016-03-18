package muni

import "fmt"

//Bus a muni bus
type Bus struct {
	Plate string
	Color string
	Brand string
}

//Move the bus starts moving
func (b *Bus) Move() error {
	d, err := b.Description()
	if err != nil {
		return err
	}
	fmt.Printf(d)
	return nil
}

//Description the bus description
func (b *Bus) Description() (string, error) {
	return fmt.Sprintf("Bus moving Placa: %s Marca: %s Color: %s...\n", b.Plate, b.Brand, b.Color), nil
}
