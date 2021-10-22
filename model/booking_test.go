package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBooking(t *testing.T) {
	booking := Booking{
		Vehicle:   &Vehicle{},
		Branch:    &Branch{},
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}
	assert.NotNil(t, booking)
}