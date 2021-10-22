package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVehicle_GetId(t *testing.T) {
	vehicle := &Vehicle{
		Id:     "1",
		Type:   "Sedan",
		Booked: false,
	}

	assert.Equal(t, "1", vehicle.Id)
}

func TestVehicle_GetType(t *testing.T) {
	vehicle := &Vehicle{
		Id:     "1",
		Type:   "Sedan",
		Booked: false,
	}

	assert.Equal(t, "Sedan", vehicle.Type)
}

func TestVehicle_IsBooked(t *testing.T) {
	vehicle := &Vehicle{
		Id:     "1",
		Type:   "Sedan",
		Booked: false,
	}

	assert.Equal(t, false, vehicle.IsBooked())
}