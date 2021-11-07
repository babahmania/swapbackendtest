package entity

import (
	"html"
	"strings"
	"time"
)

type Flight struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FlightNumber  string `gorm:"size:15;not null;unique" json:"flight_number" validate:"required,lte=15"`
	AirlineID     uint   `gorm:"column:airline_id; index:idx_flight_airline_id; not null;" json:"airline_id" validate:"required,gte=1"`
	OriginID      uint   `gorm:"column:origin_id; index:idx_flight_origin_id; not null;" json:"origin_id" validate:"required,gte=1"`
	DestinationID uint   `gorm:"column:destination_id; index:idx_flight_destination_id; not null;" json:"destination_id" validate:"required,gte=1"`

	DepartDatetime  string `gorm:"timestamp; not null" json:"depart_datetime,omitempty"`
	ArrivalDatetime string `gorm:"timestamp; not null" json:"arrival_datetime,omitempty"`
	Duration        string `gorm:"size:25;not null;" json:"duration" validate:"required,lte=25"`

	Price          uint   `gorm:"not null;" json:"price" validate:"required,gte=1"`
	SeatsAvailable uint   `gorm:"not null;" json:"seats_available" validate:"required,gte=1"`
	QtyTransit     uint   `gorm:"not null;default:0; index:idx_flight_qty_transit;" json:"qty_transit" validate:"gte=0"`
	FlightStatus   string `gorm:"column:flight_status; size:25;not null; index:idx_flight_status;" json:"flight_status" validate:"required,lte=25"`

	UserIDSubmit uint64 `gorm:"column:user_id_submit; index:idx_flight_user_id_submit; not null;" json:"user_id_submit" validate:"required,gte=1"`
	UserIDUpdate uint64 `gorm:"column:user_id_update; index:idx_flight_user_id_update; default 0;" json:"user_id_update"`
	UserIDDelete uint64 `gorm:"column:user_id_delete; index:idx_flight_user_id_delete; default 0;" json:"user_id_delete"`

	//FlightDetail
	TransitFirst  string `gorm:"size:100;" json:"transit_first"`
	TransitSecond string `gorm:"size:100;" json:"transit_second"`
	TransitThird  string `gorm:"size:100;" json:"transit_third"`

	IsEconomy                    string `gorm:"column:is_economy; size:1;not null;default:'1'" json:"is_economy" validate:"required,lte=1"`
	SeatsAvailableEconomy        uint   `gorm:"column:seats_available_economy; not null;default:0;" json:"seats_available_economy" validate:"gte=0"`
	IsPremiumEconomy             string `gorm:"column:is_premium_economy; size:1;not null;default:'1'" json:"is_premium_economy" validate:"required,lte=1"`
	SeatsAvailablePremiumEconomy uint   `gorm:"column:seats_available_premium_economy; not null;default:0;" json:"seats_available_premium_economy" validate:"gte=0"`
	IsBusiness                   string `gorm:"column:is_business; size:1;not null;default:'1'" json:"is_business" validate:"required,lte=1"`
	SeatsAvailableBusiness       uint   `gorm:"column:seats_available_business; not null;default:0;" json:"seats_available_business" validate:"gte=0"`
	IsFirstClass                 string `gorm:"column:is_first_class; size:1;not null;default:'0'" json:"is_first_class" validate:"required,lte=1"`
	SeatsAvailableFirstClass     uint   `gorm:"column:seats_available_first_class; not null;default:0;" json:"seats_available_first_class" validate:"gte=0"`

	QtyBaggage      uint   `gorm:"column:qty_baggage; not null;default:0;" json:"qty_baggage" validate:"required,gte=0"`
	QtyCabin        uint   `gorm:"column:qty_cabin; not null;default:0;" json:"qty_cabin" validate:"required,gte=0"`
	AircrafID       uint   `gorm:"column:aircraft_id; not null;" json:"aircraft_id" validate:"required,gte=1"`
	IsMeal          string `gorm:"column:is_meal; size:1;not null;default:'0'" json:"is_meal" validate:"required,lte=1"`
	IsEntertainment string `gorm:"column:is_entertainment; size:1;not null;default:'0'" json:"is_entertainment" validate:"required,lte=1"`
	IsPowerUSB      string `gorm:"column:is_power_usb; size:1;not null;default:'0'" json:"is_power_usb" validate:"required,lte=1"`

	IsActive  string     `gorm:"size:1;not null;default:'1'; index:idx_flight_active_status;" json:"is_active" validate:"required,lte=1"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type FlightView struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FlightNumber     string `gorm:"size:15;not null;unique" json:"flight_number" validate:"required,lte=15"`
	AirlineID        string `json:"airline_id"`
	AirlineName      string `json:"airline_name"`
	AirlineImageName string `json:"airline_image_name"`
	OriginName       string `json:"origin_name"`
	OriginCode       string `json:"origin_code"`
	DestinationName  string `json:"destination_name"`
	DestinationCode  string `json:"destination_code"`
	AircraftName     string `json:"aircraft_name"`

	DepartDatetime  string `gorm:"timestamp; not null" json:"depart_datetime,omitempty"`
	ArrivalDatetime string `gorm:"timestamp; not null" json:"arrival_datetime,omitempty"`
	Duration        string `gorm:"size:25;not null;" json:"duration" validate:"required,lte=25"`

	Price          uint   `gorm:"not null;" json:"price" validate:"required,gte=1"`
	SeatsAvailable uint   `gorm:"not null;" json:"seats_available" validate:"required,gte=1"`
	QtyTransit     uint   `gorm:"not null;default:0; index:idx_flight_qty_transit;" json:"qty_transit" validate:"gte=0"`
	FlightStatus   string `gorm:"column:flight_status; size:25;not null; index:idx_flight_status;" json:"flight_status" validate:"required,lte=25"`

	//FlightDetail
	TransitFirst  string `gorm:"size:100;" json:"transit_first"`
	TransitSecond string `gorm:"size:100;" json:"transit_second"`
	TransitThird  string `gorm:"size:100;" json:"transit_third"`

	IsEconomy                    string `gorm:"column:is_economy; size:1;not null;default:'1'" json:"is_economy" validate:"required,lte=1"`
	SeatsAvailableEconomy        uint   `gorm:"column:seats_available_economy; not null;default:0;" json:"seats_available_economy" validate:"gte=0"`
	IsPremiumEconomy             string `gorm:"column:is_premium_economy; size:1;not null;default:'1'" json:"is_premium_economy" validate:"required,lte=1"`
	SeatsAvailablePremiumEconomy uint   `gorm:"column:seats_available_premium_economy; not null;default:0;" json:"seats_available_premium_economy" validate:"gte=0"`
	IsBusiness                   string `gorm:"column:is_business; size:1;not null;default:'1'" json:"is_business" validate:"required,lte=1"`
	SeatsAvailableBusiness       uint   `gorm:"column:seats_available_business; not null;default:0;" json:"seats_available_business" validate:"gte=0"`
	IsFirstClass                 string `gorm:"column:is_first_class; size:1;not null;default:'0'" json:"is_first_class" validate:"required,lte=1"`
	SeatsAvailableFirstClass     uint   `gorm:"column:seats_available_first_class; not null;default:0;" json:"seats_available_first_class" validate:"gte=0"`

	QtyBaggage uint `gorm:"column:qty_baggage; not null;default:0;" json:"qty_baggage" validate:"required,gte=0"`
	QtyCabin   uint `gorm:"column:qty_cabin; not null;default:0;" json:"qty_cabin" validate:"required,gte=0"`

	IsMeal          string `gorm:"column:is_meal; size:1;not null;default:'0'" json:"is_meal" validate:"required,lte=1"`
	IsEntertainment string `gorm:"column:is_entertainment; size:1;not null;default:'0'" json:"is_entertainment" validate:"required,lte=1"`
	IsPowerUSB      string `gorm:"column:is_power_usb; size:1;not null;default:'0'" json:"is_power_usb" validate:"required,lte=1"`

	IsActive string `gorm:"size:1;not null;default:'1'; index:idx_flight_active_status;" json:"is_active" validate:"required,lte=1"`
}

//BeforeSave is a gorm hook
func (a *Flight) BeforeSave() {
	a.FlightNumber = html.EscapeString(strings.TrimSpace(a.FlightNumber))
}

type Flights []Flight

func (a *Flight) Prepare() {
	a.FlightNumber = html.EscapeString(strings.TrimSpace(a.FlightNumber))
	a.Duration = html.EscapeString(strings.TrimSpace(a.Duration))
	a.IsActive = html.EscapeString(strings.TrimSpace(a.IsActive))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Flight) Validate(action string) string {
	switch strings.ToLower(action) {
	case "update":
		if a.FlightNumber == "" {
			return "flight number required"
		}
		if a.IsActive == "" {
			return "status active required"
		}

	default:
		if a.FlightNumber == "" {
			return "flight number required"
		}
		if a.Duration == "" {
			return "duration required"
		}
		if a.IsActive == "" {
			return "status active required"
		}
	}
	return ""
}
