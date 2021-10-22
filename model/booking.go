package model

import "time"

type Booking struct {
	Vehicle   *Vehicle
	Branch    *Branch
	StartTime time.Time
	EndTime   time.Time
}
