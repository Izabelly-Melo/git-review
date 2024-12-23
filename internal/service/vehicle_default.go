package service

import (
	"app/internal"
	"errors"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) UpdateFuel(id int, fuelType string) (v internal.Vehicle, err error) {

	vehicles, _ := s.rp.FindAll()
	exists := false
	for _, vehicle := range vehicles {
		if vehicle.Id == id {
			v, _ = s.rp.UpdateFuel(id, fuelType)
			exists = true
		}
	}

	if !exists {
		err = errors.New("not Found")
	}
	return v, err
}
