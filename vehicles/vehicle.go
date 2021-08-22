package vehicles

import (
	"github.com/burxtx/car/vehicles/utils"
)

type VehicleID string

// Vehicle is the central of post domain model
// play as aggragate root
type Vehicle struct {
	ID      VehicleID
	Name    string
	NameCN  string
	Brand   string
	BrandCN string
	Model   Model
}

type Model struct {
	Name string
	Year string
}

// NewUser creates a new user
func NewVehicle(name, brand string, model Model) *Vehicle {
	return &Vehicle{
		ID:    VehicleID(utils.NewVehicleID()),
		Model: model,
	}
}

func NewModel(name, year string) Model {
	return Model{
		Name: name,
		Year: year,
	}
}

// Repository provides access to a store.
type VehicleRepository interface {
	Create(vehicle *Vehicle) error
	Find(VehicleID string) (*Vehicle, error)
	List() []*Vehicle
}
