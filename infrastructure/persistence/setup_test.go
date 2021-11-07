package persistence

import (
	"fmt"
	"log"
	"os"
	"swapbackendtest/domain/entity"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func DBConn() (*gorm.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		//var err error
		err := godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return LocalDatabase()
	}
	return CIBuild()
}

//Circle CI DB
func CIBuild() (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "192.168.1.212", "3308", "flight-app-fiber-test")
	conn, err := gorm.Open("mysql", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, nil
}

//Local DB
func LocalDatabase() (*gorm.DB, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "192.168.1.212", "3308", "flight-app-fiber-test")
	dbdriver := "mysql"
	conn, err := gorm.Open(dbdriver, DBURL)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbdriver)
	}

	err = conn.DropTableIfExists(&entity.User{}, &entity.Airport{}, &entity.Flight{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		entity.User{},
		entity.Airport{},
		entity.Flight{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func seedUser(db *gorm.DB) (*entity.User, error) {
	user := &entity.User{
		ID:         1,
		FullName:   "babah",
		UserStatus: "1",
		Email:      "babahmania@gmail.com",
		Password:   "password",
		DeletedAt:  nil,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func seedUsers(db *gorm.DB) ([]entity.User, error) {
	users := []entity.User{
		{
			ID:         1,
			FullName:   "babah",
			UserStatus: "1",
			Email:      "babahmania@gmail.com",
			Password:   "password",
			DeletedAt:  nil,
		},
		{
			ID:         2,
			FullName:   "kurnia",
			UserStatus: "0",
			Email:      "maniababah@gmail.com",
			Password:   "password",
			DeletedAt:  nil,
		},
	}
	for _, v := range users {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
