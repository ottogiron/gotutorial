package workers

import "github.com/ottogiron/gotutorial/muni"

type Driver struct {
}

func (d *Driver) Drive(vehicle muni.Vehicle) {
	vehicle.Move()
}
