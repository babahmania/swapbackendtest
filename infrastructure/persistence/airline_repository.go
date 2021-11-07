package persistence

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
	"swapbackendtest/infrastructure/validator"

	"github.com/jinzhu/gorm"
)

type AirlineRepo struct {
	db *gorm.DB
}

func NewAirlineRepository(db *gorm.DB) *AirlineRepo {
	return &AirlineRepo{db}
}

//AirlineRepo implements the repository.AirlineRepository interface
var _ repository.AirlineRepository = &AirlineRepo{}

func (r *AirlineRepo) SaveAirline(airline *entity.Airline) (*entity.Airline, error) {
	// Create a new validator for a Airline model.
	validate := validator.NewValidator()
	if err := validate.Struct(airline); err != nil {
		// Return, if some fields are not valid.
		msgNotValid := validator.ValidatorErrors(err)
		b := new(bytes.Buffer)
		for _, value := range msgNotValid {
			fmt.Fprintf(b, "%s. ", value)
		}
		msgError := b.String()
		return nil, errors.New(msgError)
	}
	err := r.db.Debug().Create(&airline).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("name already taken")
		}
		return nil, errors.New("database error")
	}
	return airline, nil
}

func (r *AirlineRepo) GetAirline(id uint64) (*entity.Airline, error) {
	var airline entity.Airline
	err := r.db.Debug().Where("id = ?", id).Take(&airline).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("airline not found")
	}
	return &airline, nil
}

func (r *AirlineRepo) GetAirlines() ([]entity.Airline, error) {
	var airlines []entity.Airline
	err := r.db.Debug().Order("name").Find(&airlines).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("airline not found")
	}
	return airlines, nil
}
