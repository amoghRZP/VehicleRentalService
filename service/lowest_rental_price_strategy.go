package service

import "VechicleRentalService/model"

type LowestRentalPriceStrategy struct {
}

func (lrp *LowestRentalPriceStrategy) FindVehicle(vehicleType string, bookings []*model.Booking,
	branches []*model.Branch) (*model.Branch, *model.Vehicle) {
	var minPrice int
	var resultVehicle *model.Vehicle
	var resultBranch *model.Branch
	for _, branch := range branches {
		vehicle := branch.FindLowestPricedVehicleAvailable(vehicleType)
		if vehicle != nil {
			if minPrice == 0 {
				minPrice = branch.Price[vehicleType]
				resultVehicle = vehicle
				resultBranch = branch
			} else {
				if minPrice > branch.Price[vehicleType] {
					minPrice = branch.Price[vehicleType]
					resultVehicle = vehicle
					resultBranch = branch
				}
			}
		}
	}

	return resultBranch, resultVehicle
}
