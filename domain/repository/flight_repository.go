package repository

import (
	"swapbackendtest/domain/entity"
)

type FlightRepository interface {
	SaveFlight(*entity.Flight) (*entity.Flight, error)
	GetFlight(uint64) (*entity.Flight, error)
	GetFlights() ([]entity.Flight, error)
	GetFlightLocation(uint64, uint64, int) ([]entity.Flight, error)
	GetFlightLocationFind(uint64, uint64, int) ([]entity.Flight, error)
}
