package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBranch_AddVehicle(t *testing.T) {
	branch := &Branch{
		Name:     "Vasant Vihar",
		Price:    make(map[string]int),
	}
	branch.AddVehicle("1234", "Sedan")
	assert.Equal(t, 1, len(branch.Vehicles))
}
