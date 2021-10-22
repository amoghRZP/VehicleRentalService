package model

const (
	VehicleTypeSuv       = "Suv"
	VehicleTypeSedan     = "Sedan"
	VehicleTypeHatchback = "Hatchback"
)

type IVehicle interface {
	GetType() string
	GetId() string
	IsBooked() bool
}

type Vehicle struct {
	Id   string
	Type string
	Booked bool
}

func (s *Vehicle) GetType() string {
	return s.Type
}

func (s *Vehicle) GetId() string {
	return s.Id
}

func (s *Vehicle) IsBooked() bool {
	return s.Booked
}
