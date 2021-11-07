package repository

import (
	"swapbackendtest/domain/entity"
)

type AirlineRepository interface {
	SaveAirline(*entity.Airline) (*entity.Airline, error)
	GetAirline(uint64) (*entity.Airline, error)
	GetAirlines() ([]entity.Airline, error)
}
