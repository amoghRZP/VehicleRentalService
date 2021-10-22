package main

import (
	"VechicleRentalService/service"
	"time"
)

func main() {
	rentalService := &service.VehicleRental{}
	rentalService.SetBookingStrategy("LowestRentalPriceStrategy")

	rentalService.AddBranch("Vasanth Vihar")
	rentalService.AddBranch("Cyber City")

	rentalService.AllocatePrice("Vasanth Vihar", "Sedan", 100)
	rentalService.AllocatePrice("Vasanth Vihar", "Hatchback", 80)
	rentalService.AllocatePrice("Cyber City", "Sedan", 200)
	rentalService.AllocatePrice("Cyber City", "Hatchback", 50)

	rentalService.AddVehicle("DL 01 MR 9310", "Sedan", "Vasanth Vihar")
	rentalService.AddVehicle("DL 01 MR 9311", "Sedan", "Cyber City")
	rentalService.AddVehicle("DL 01 MR 9312", "Hatchback", "Cyber City")

	rentalService.BookVehicle("Sedan", time.Now(), time.Now())
	rentalService.BookVehicle("Sedan", time.Now(), time.Now())
	rentalService.BookVehicle("Sedan", time.Now(), time.Now())

}
