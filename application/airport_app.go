package application

import (
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
)

type airportApp struct {
	ap repository.AirportRepository
}

//AirportApp implements the AirportAppInterface
var _ AirportAppInterface = &airportApp{}

type AirportAppInterface interface {
	SaveAirport(*entity.Airport) (*entity.Airport, error)
	GetAirports() ([]entity.Airport, error)
	GetAirport(uint64) (*entity.Airport, error)
}

func (a *airportApp) SaveAirport(airport *entity.Airport) (*entity.Airport, error) {
	return a.ap.SaveAirport(airport)
}

func (a *airportApp) GetAirport(airportId uint64) (*entity.Airport, error) {
	return a.ap.GetAirport(airportId)
}

func (a *airportApp) GetAirports() ([]entity.Airport, error) {
	return a.ap.GetAirports()
}
