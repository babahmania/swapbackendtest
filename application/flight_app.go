package application

import (
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
)

type flightApp struct {
	fl repository.FlightRepository
}

//FlightApp implements the FlightAppInterface
var _ FlightAppInterface = &flightApp{}

type FlightAppInterface interface {
	SaveFlight(*entity.Flight) (*entity.Flight, error)
	GetFlights() ([]entity.Flight, error)
	GetFlight(uint64) (*entity.Flight, error)
	GetFlightLocation(uint64, uint64, int) ([]entity.Flight, error)
	GetFlightLocationFind(uint64, uint64, int) ([]entity.Flight, error)
}

func (a *flightApp) SaveFlight(flight *entity.Flight) (*entity.Flight, error) {
	return a.fl.SaveFlight(flight)
}

func (a *flightApp) GetFlight(flightId uint64) (*entity.Flight, error) {
	return a.fl.GetFlight(flightId)
}

func (a *flightApp) GetFlights() ([]entity.Flight, error) {
	return a.fl.GetFlights()
}
func (a *flightApp) GetFlightLocation(originId uint64, destinationId uint64, qtyTransit int) ([]entity.Flight, error) {
	return a.fl.GetFlightLocation(originId, destinationId, qtyTransit)
}
func (a *flightApp) GetFlightLocationFind(originId uint64, destinationId uint64, qtyTransit int) ([]entity.Flight, error) {
	return a.fl.GetFlightLocationFind(originId, destinationId, qtyTransit)
}
