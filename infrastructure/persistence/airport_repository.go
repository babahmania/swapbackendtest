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

type AirportRepo struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) *AirportRepo {
	return &AirportRepo{db}
}

//AirportRepo implements the repository.AirportRepository interface
var _ repository.AirportRepository = &AirportRepo{}

func (r *AirportRepo) SaveAirport(airport *entity.Airport) (*entity.Airport, error) {
	// Create a new validator for a Airport model.
	validate := validator.NewValidator()
	if err := validate.Struct(airport); err != nil {
		// Return, if some fields are not valid.
		msgNotValid := validator.ValidatorErrors(err)
		b := new(bytes.Buffer)
		for _, value := range msgNotValid {
			fmt.Fprintf(b, "%s. ", value)
		}
		msgError := b.String()
		return nil, errors.New(msgError)
	}
	err := r.db.Debug().Create(&airport).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("name already taken")
		}
		return nil, errors.New("database error")
	}
	return airport, nil
}

func (r *AirportRepo) GetAirport(id uint64) (*entity.Airport, error) {
	var airport entity.Airport
	err := r.db.Debug().Where("id = ?", id).Take(&airport).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("airport not found")
	}
	return &airport, nil
}

func (r *AirportRepo) GetAirports() ([]entity.Airport, error) {
	var airports []entity.Airport
	err := r.db.Debug().Order("city").Find(&airports).Error
	if err != nil {
		return nil, err
	}
	//if gorm.IsRecordNotFoundError(err) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("airport not found")
	}
	return airports, nil
}
