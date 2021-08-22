package manage

import (
	"context"

	"github.com/burxtx/car/vehicles"
)

// Service is the interface that provides vehicles related methods.
type Service interface {
	Create(ctx context.Context, name, brand, modelName, year string) (*vehicles.Vehicle, error)
}

type service struct {
	repo vehicles.VehicleRepository
}

func New(repo vehicles.VehicleRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, name, brand, modelName, year string) (*vehicles.Vehicle, error) {
	vm := vehicles.NewModel(modelName, year)
	v := vehicles.NewVehicle(name, brand, vm)
	err := s.repo.Create(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Vehicle is a real model for view
type Vehicle struct {
	VehicleID string
	Name      string
	Brand     string
	ModelName string
}
