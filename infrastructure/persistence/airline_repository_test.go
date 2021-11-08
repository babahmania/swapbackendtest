package persistence

import (
	"swapbackendtest/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAirline_Failure(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var data = entity.Airline{}
	data.Name = "Air Asia"
	data.ImageName = "Air Asia"
	data.IsActive = "1"

	repo := NewAirlineRepository(conn)
	u, saveErr := repo.SaveAirline(&data)
	dbMsg := "name already taken"
	assert.Nil(t, u)
	assert.EqualError(t, saveErr, dbMsg)
}

func TestGetAirline_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewAirlineRepository(conn)
	u, getErr := repo.GetAirline(1)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Name, "Air Asia")
}

func TestGetAirlines_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewAirlineRepository(conn)
	datas, getErr := repo.GetAirlines()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(datas), 5)
}

func TestGetAirlineByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var id uint64 = 1
	repo := NewAirlineRepository(conn)
	u, getErr := repo.GetAirline(id)

	assert.Nil(t, getErr)
	assert.NotEqual(t, u.ID, 1)
}
