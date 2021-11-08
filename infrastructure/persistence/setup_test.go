package persistence

import (
	"fmt"
	"log"
	"os"

	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
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
	//DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "192.168.1.212", "3308", "flight-app-fiber-test")
	//conn, err := gorm.Open("mysql", DBURL)
	dsn := fmt.Sprintf("%s:%s@tcp(192.168.1.212:3308)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "flight-app-fiber-test")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, nil
}

//Local DB
func LocalDatabase() (*gorm.DB, error) {
	/*
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "192.168.1.212", "3308", "flight-app-fiber-test")
		dbdriver := "mysql"
		conn, err := gorm.Open(dbdriver, DBURL)
	*/
	dsn := fmt.Sprintf("%s:%s@tcp(192.168.1.212:3308)/%s?charset=utf8&parseTime=True&loc=Local", "swap", "Sw4p-1443", "flight-app-fiber-test")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", "database unit testing")
	}

	/*
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
	*/
	return conn, nil
}
