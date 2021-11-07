package application

import (
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
)

type airlineApp struct {
	al repository.AirlineRepository
}

//AirlineApp implements the AirlineAppInterface
var _ AirlineAppInterface = &airlineApp{}

type AirlineAppInterface interface {
	SaveAirline(*entity.Airline) (*entity.Airline, error)
	GetAirlines() ([]entity.Airline, error)
	GetAirline(uint64) (*entity.Airline, error)
}

func (a *airlineApp) SaveAirline(airline *entity.Airline) (*entity.Airline, error) {
	return a.al.SaveAirline(airline)
}

func (a *airlineApp) GetAirline(airlineId uint64) (*entity.Airline, error) {
	return a.al.GetAirline(airlineId)
}

func (a *airlineApp) GetAirlines() ([]entity.Airline, error) {
	return a.al.GetAirlines()
}
