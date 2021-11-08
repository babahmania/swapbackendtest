package persistence

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
	"swapbackendtest/infrastructure/validator"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type FlightRepo struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) *FlightRepo {
	return &FlightRepo{db}
}

//FlightRepo implements the repository.FlightRepository interface
var _ repository.FlightRepository = &FlightRepo{}

func (r *FlightRepo) SaveFlight(flight *entity.Flight) (*entity.Flight, error) {
	// Create a new validator for a Flight model.
	validate := validator.NewValidator()
	if err := validate.Struct(flight); err != nil {
		// Return, if some fields are not valid.
		msgNotValid := validator.ValidatorErrors(err)
		b := new(bytes.Buffer)
		for _, value := range msgNotValid {
			fmt.Fprintf(b, "%s. ", value)
		}
		msgError := b.String()
		return nil, errors.New(msgError)
	}

	err := r.db.Debug().Create(&flight).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("name already taken")
		}
		return nil, errors.New("database error")
	}
	return flight, nil
}

func (r *FlightRepo) GetFlight(id uint64) (*entity.Flight, error) {
	var flight entity.Flight
	err := r.db.Debug().Where("id = ?", id).Take(&flight).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//if gorm.ErrRecordNotFound(err) {
		return nil, errors.New("flight not found")
	}
	return &flight, nil
}

func (r *FlightRepo) GetFlights() ([]entity.Flight, error) {
	var flights []entity.Flight
	err := r.db.Debug().Find(&flights).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("flight not found")
	}
	return flights, nil
}

func (r *FlightRepo) GetFlightLocation(originId uint64, destinationId uint64, qtyTransit int) ([]entity.Flight, error) {
	var flights []entity.Flight
	var err error
	if qtyTransit == 100 {
		err = r.db.Debug().Where("origin_id = ? and destination_id = ?", originId, destinationId).Find(&flights).Error
	} else {
		err = r.db.Debug().Where("origin_id = ? and destination_id = ? and qty_transit=?", originId, destinationId, qtyTransit).Find(&flights).Error
	}
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("flight not found")
	}
	return flights, nil
}
func (r *FlightRepo) GetFlightLocationFind(originId uint64, destinationId uint64, qtyTransit int) ([]entity.Flight, error) {
	var flights []entity.Flight
	var err error
	if qtyTransit == 100 {
		err = r.db.Raw("select * from view_flights where origin_id = ? and destination_id = ?", originId, destinationId).Scan(&flights).Error
		//err = r.db.Debug().Where("origin_id = ? and destination_id = ?", originId, destinationId).Find(&flights).Error
		//err = r.db.Debug().Where("origin_id = ? and destination_id = ?", originId, destinationId).Find(&flights).Error
	} else {
		err = r.db.Raw("select * from view_flights where origin_id = ? and destination_id = ? and qty_transit=?", originId, destinationId, qtyTransit).Scan(&flights).Error
		//err = r.db.Debug().Where("origin_id = ? and destination_id = ? and qty_transit=?", originId, destinationId, qtyTransit).Find(&flights).Error
	}

	//var users []User
	//db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)

	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("flight not found")
	}
	return flights, nil
}
