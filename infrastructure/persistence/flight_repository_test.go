package persistence

import (
	"swapbackendtest/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveFlight_Failure(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var data = entity.Flight{}
	data.FlightNumber = "CKG-001"
	data.AirlineID = 1
	data.AircrafID = 1
	data.OriginID = 1
	data.DestinationID = 3
	data.DepartDatetime = "2021-11-08 06:00:00"
	data.ArrivalDatetime = "2021-11-08 07:00:00"
	data.Duration = "1h 55m"
	data.Price = 1000000
	data.SeatsAvailable = 100
	data.QtyTransit = 0
	data.FlightStatus = "open booking"
	data.TransitFirst = ""
	data.IsMeal = "0"
	data.IsEntertainment = "0"
	data.IsPowerUSB = "0"
	data.QtyBaggage = 20
	data.QtyCabin = 10
	data.IsEconomy = "1"
	data.SeatsAvailableEconomy = 100
	data.IsPremiumEconomy = "0"
	data.SeatsAvailableEconomy = 0
	data.IsBusiness = "0"
	data.SeatsAvailableBusiness = 0
	data.IsFirstClass = "0"
	data.SeatsAvailableFirstClass = 0
	data.UserIDSubmit = 1
	data.IsActive = "1"

	repo := NewFlightRepository(conn)
	u, saveErr := repo.SaveFlight(&data)
	dbMsg := "name already taken"
	assert.Nil(t, u)
	assert.EqualError(t, saveErr, dbMsg)
}

func TestGetFlight_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewFlightRepository(conn)
	u, getErr := repo.GetFlight(1)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.FlightNumber, "CKG-001")
}

func TestGetFlights_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewFlightRepository(conn)
	datas, getErr := repo.GetFlights()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(datas), 6)
}

func TestGetFlightByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var id uint64 = 1
	repo := NewFlightRepository(conn)
	u, getErr := repo.GetFlight(id)

	assert.Nil(t, getErr)
	assert.NotEqual(t, u.ID, 1)
}
