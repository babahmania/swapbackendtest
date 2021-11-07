package repository

import (
	"swapbackendtest/domain/entity"
)

type AirportRepository interface {
	SaveAirport(*entity.Airport) (*entity.Airport, error)
	GetAirport(uint64) (*entity.Airport, error)
	GetAirports() ([]entity.Airport, error)
}
