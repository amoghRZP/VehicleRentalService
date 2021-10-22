package model

type Branch struct {
	Name     string
	Vehicles []*Vehicle
	Price    map[string]int
}

func (b *Branch) AddVehicle(vehicleId string, vehicleType string) {
	vehicle := &Vehicle{
		Id:   vehicleId,
		Type: vehicleType,
	}
	b.Vehicles = append(b.Vehicles, vehicle)
}

func (b *Branch) FindLowestPricedVehicleAvailable(vehicleType string) *Vehicle{
	for _, vehicle := range b.Vehicles {
		if vehicle.GetType() == vehicleType && !vehicle.IsBooked() {
			return vehicle
		}
	}
	return nil
}

