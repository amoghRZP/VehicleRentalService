package service

import "VechicleRentalService/model"

type IBookingStrategy interface {
	FindVehicle(string, []*model.Booking, []*model.Branch) (*model.Branch, *model.Vehicle)
}
