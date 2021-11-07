package persistence

import (
	"errors"
	"fmt"
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
	"time"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	User    repository.UserRepository
	Airport repository.AirportRepository
	Airline repository.AirlineRepository
	Flight  repository.FlightRepository
	db      *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User:    NewUserRepository(db),
		Airport: NewAirportRepository(db),
		Airline: NewAirlineRepository(db),
		Flight:  NewFlightRepository(db),
		db:      db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//FlightView
func (s *Repositories) GetFlightLocationFind(originId uint64, destinationId uint64, qtyTransit int, dateFlight time.Time,
	sortFlight string, departureTime string, classFlight string, airlineFlight int) ([]entity.FlightView, error) {
	var flights []entity.FlightView
	var err error
	orderBy := "depart_datetime"
	switch sortFlight {
	case "1":
		orderBy = "arrival_datetime"
	case "2":
		orderBy = "airline_id"
	case "40":
		orderBy = "price"
	case "41":
		orderBy = "price DESC"
	}

	depDate := fmt.Sprint(dateFlight)
	depDate = depDate[0:10]
	sql := "select * from view_flights where date(depart_datetime)='" + depDate + "' and origin_id = " + fmt.Sprint(originId) + " and destination_id = " + fmt.Sprint(destinationId)
	if qtyTransit != 100 {
		sql += " and qty_transit= " + fmt.Sprint(qtyTransit)
	}

	depTime := ""
	switch departureTime {
	case "0006":
		depTime = " and depart_datetime < '" + depDate + " 06:00:00'"
	case "0612":
		depTime = " and depart_datetime BETWEEN '" + depDate + " 00:00:00' and '" + depDate + " 05:59:59'"
	case "1218":
		depTime = " and depart_datetime BETWEEN '" + depDate + " 12:00:00' and '" + depDate + " 17:59:59'"
	case "1824":
		depTime = " and depart_datetime >= '" + depDate + " 18:00:00'"
	}
	sql += depTime

	flightClass := ""
	switch classFlight {
	case "1": //economy
		flightClass = " and is_economy ='1'"
	case "2":
		flightClass = " and is_premium_economy ='1'"
	case "3":
		flightClass = " and is_business ='1'"
	case "4":
		flightClass = " and is_first_class ='1'"
	}
	sql += flightClass

	if airlineFlight != 0 {
		sql += " and airline_id =" + fmt.Sprint(airlineFlight)
	}

	//fmt.Println(airlineFlight)
	/*
		fmt.Println(flightClass)
		fmt.Println(dateFlight)
		fmt.Println(orderBy)
		fmt.Println(departureTime)
		fmt.Println(depTime)
	*/
	fmt.Println(sql)
	err = s.db.Raw(sql).Order(orderBy).Scan(&flights).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("flight not found")
	}
	return flights, nil
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	results := s.db.AutoMigrate(&entity.User{},
		&entity.Airport{}, &entity.Airline{}, &entity.Aircraft{},
		&entity.Flight{},
	).Error
	if results == nil {
		s.db.Model(&entity.Flight{}).AddForeignKey("airline_id", "airlines(id)", "RESTRICT", "RESTRICT")
		s.db.Model(&entity.Flight{}).AddForeignKey("origin_id", "airports(id)", "RESTRICT", "RESTRICT")
		s.db.Model(&entity.Flight{}).AddForeignKey("destination_id", "airports(id)", "RESTRICT", "RESTRICT")
		s.db.Model(&entity.Flight{}).AddForeignKey("aircraft_id", "aircrafts(id)", "RESTRICT", "RESTRICT")
		s.db.Model(&entity.Flight{}).AddForeignKey("user_id_submit", "users(id)", "RESTRICT", "RESTRICT")

		//CreateView(name string, option ViewOption) error
		//db.Migrator().CreateTable(&User{})
		//s.db.Migrator()
		/*
					//"view_flights", "SELECT `flights`.`id`,
			    `flights`.`flight_number`,
			    `flights`.`airline_id`,
			    `airlines`.`name` as airline_name,
			    `airlines`.`image_name` as airline_image_name,
			    `flights`.`origin_id`,
			    origin_airport.`name` as origin_name,
			    origin_airport.`code` as origin_code,
			    `flights`.`destination_id`,
			    destination_airport.`name` as destination_name,
			    destination_airport.`code` as destination_code,
			    `flights`.`aircraft_id`,
			    `aircrafts`.`name` as aircraft_name,
			    `flights`.`depart_datetime`,
			    `flights`.`arrival_datetime`,
			    `flights`.`duration`,
			    `flights`.`price`,
			    `flights`.`seats_available`,
			    `flights`.`qty_transit`,
			    `flights`.`flight_status`,
			    `flights`.`user_id_submit`,
			    `flights`.`user_id_update`,
			    `flights`.`user_id_delete`,
			    `flights`.`transit_first`,
			    `flights`.`transit_second`,
			    `flights`.`transit_third`,
			    `flights`.`is_economy`,
			    `flights`.`seats_available_economy`,
			    `flights`.`is_premium_economy`,
			    `flights`.`seats_available_premium_economy`,
			    `flights`.`is_business`,
			    `flights`.`seats_available_business`,
			    `flights`.`is_first_class`,
			    `flights`.`seats_available_first_class`,
			    `flights`.`qty_baggage`,
			    `flights`.`qty_cabin`,
			    `flights`.`is_meal`,
			    `flights`.`is_entertainment`,
			    `flights`.`is_power_usb`,
			    `flights`.`is_active`
			FROM `flight-app-fiber`.`flights`, `flight-app-fiber`.`airlines`,
			`flight-app-fiber`.`airports` as origin_airport, `flight-app-fiber`.`airports` as destination_airport,
			`flight-app-fiber`.`aircrafts`
			where `flights`.`airline_id`=`airlines`.`id` and
			`flights`.`origin_id`=origin_airport.`id` and
			`flights`.`destination_id`=destination_airport.`id` and
			`flights`.`aircraft_id`=`aircrafts`.`id`
			order by `flights`.`depart_datetime`"
		*/
	}
	return results
}
