package persistence

import (
	"errors"
	"fmt"
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
	"time"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	//dsn := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)
	//DBURL := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/%s?charset=utf8&parseTime=True&loc=Asia/Jakarta", DbUser, DbPassword, DbHost, DbPort, DbName)
	//db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	//dsn := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)
	//db, err := sql.Open(Dbdriver, DBURL)
	//DBURL := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Asia/Jakarta", DbUser, DbPassword, DbName)
	//DBURL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Asia/Jakarta", DbUser, DbPassword, DbName)
	//DBURL := fmt.Sprintf("%s:%s@tcp(172.30.32.1:3306)/%s?charset=utf8&parseTime=True&loc=Asia/Jakarta", DbUser, DbPassword, DbName)
	//DBURL := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Asia/Jakarta", DbUser, DbPassword, DbName)
	//db, err := gorm.Open(Dbdriver, DBURL)

	//dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbName,)
	//for docker
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbName)
	/*
		//for local
		dsn := fmt.Sprintf("%s:%s@tcp(192.168.1.212:3308)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbName)
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	//db.LogMode(true)

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
	//return s.db.Close
	return nil
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
	//fmt.Println(sql)
	err = s.db.Raw(sql).Order(orderBy).Scan(&flights).Error
	if err != nil {
		return nil, err
	}
	//if gorm.IsRecordNotFoundError(err) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("flight not found")
	}
	return flights, nil
}
