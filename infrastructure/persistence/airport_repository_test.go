package persistence

import (
	"swapbackendtest/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAirport_Failure(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var data = entity.Airport{}
	data.Name = "Jakarta Indonesia CGK Soekarno Hatta International Airport"
	data.Code = "CGK"
	data.City = "Jakarta"
	data.Country = "Indonesia"
	data.IsActive = "1"

	repo := NewAirportRepository(conn)
	u, saveErr := repo.SaveAirport(&data)
	dbMsg := "name already taken"
	assert.Nil(t, u)
	assert.EqualError(t, saveErr, dbMsg)
}

func TestGetAirport_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewAirportRepository(conn)
	u, getErr := repo.GetAirport(1)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Name, "Jakarta Indonesia CGK Soekarno Hatta International Airport")
}

func TestGetAirports_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewAirportRepository(conn)
	datas, getErr := repo.GetAirports()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(datas), 6)
}

func TestGetAirportByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var id uint64 = 1
	repo := NewAirportRepository(conn)
	u, getErr := repo.GetAirport(id)

	assert.Nil(t, getErr)
	assert.NotEqual(t, u.ID, 1)
}
