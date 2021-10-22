package service

import (
	"VechicleRentalService/model"
	"fmt"
	"strings"
	"time"
)

const LowestRentalPriceStrategyType = "LowestRentalPriceStrategy"

type VehicleRental struct {
	Branches       []*model.Branch
	Bookings       []*model.Booking
	RentalStrategy IBookingStrategy
}

func (vr *VehicleRental) SetBookingStrategy(bookingStrategyType string) {
	if bookingStrategyType == LowestRentalPriceStrategyType {
		vr.RentalStrategy = &LowestRentalPriceStrategy{}
	}
}

func (vr *VehicleRental) AddBranch(branchName string) {
	branch := &model.Branch{
		Name: branchName,
		Price: make(map[string]int),
	}
	vr.Branches = append(vr.Branches, branch)
}

func (vr *VehicleRental) AllocatePrice(branchName string, vehicleType string, price int) {
	for _, branch := range vr.Branches {
		if branch.Name == branchName {
			branch.Price[vehicleType] = price
		}
	}
}

func (vr *VehicleRental) AddVehicle(id string, vehicleType string, branchName string) {
	for _, branch := range vr.Branches {
		if branch.Name == branchName {
			branch.AddVehicle(id, vehicleType)
		}
	}
}

func (vr *VehicleRental) BookVehicle(vehicleType string, startTime time.Time, endTime time.Time) {
	branch, vehicle := vr.RentalStrategy.FindVehicle(vehicleType, vr.Bookings, vr.Branches)

	if branch == nil || vehicle == nil {
		fmt.Println("NO ", strings.ToUpper(vehicleType), " AVAILABLE")
		return
	}
	fmt.Println(vehicle.Id, " from  ", branch.Name, " booked.")

	booking := &model.Booking{
		Vehicle:   vehicle,
		Branch:    branch,
		StartTime: startTime,
		EndTime:   endTime,
	}
	vr.Bookings = append(vr.Bookings, booking)
	vehicle.Booked = true
}
