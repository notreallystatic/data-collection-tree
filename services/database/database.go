package database

import (
	"context"
	"fmt"

	model "dct/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MariaDbClient *gorm.DB

func Init() {
	fmt.Println("Connecting to DB....")
	dsn := "test:secret@tcp(maria-db:3306)/dct?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MariaDbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Error connecting to DB, error: ", err)
		panic(err)
	}

	fmt.Println("Connected to db!!!")

	var version string
	MariaDbClient.Raw("SELECT VERSION()").Scan(&version)
	fmt.Println("Database Version: ", version)

	CreateTables()
	SeedData()
}

func CreateTables() {
	fmt.Println("running migrations...")
	err := MariaDbClient.Table("country").AutoMigrate(&model.Country{})
	if err != nil {
		fmt.Printf("error creating/updating table Country, error: %v\n", err)
		panic(err)
	}
	err = MariaDbClient.Table("web_request").AutoMigrate(&model.WebRequest{})
	if err != nil {
		fmt.Printf("error creating/updating table WebRequest, error: %v\n", err)
		panic(err)
	}
	err = MariaDbClient.Table("time_spent").AutoMigrate(&model.TimeSpent{})
	if err != nil {
		fmt.Printf("error creating/updating table TimeSpent, error: %v\n", err)
		panic(err)
	}
	err = MariaDbClient.Table("device").AutoMigrate(&model.Device{})
	if err != nil {
		fmt.Printf("error creating/updating table Device, error: %v\n", err)
		panic(err)
	}
	fmt.Println("migrations ran successfully!!!")
}

func GetTransactionInstanceFromContext(ctx context.Context) *gorm.DB {
	if dbTxn := ctx.Value("dbTxn"); dbTxn != nil {
		if val, ok := dbTxn.(*gorm.DB); ok {
			return val
		}
	}
	return nil
}

func SeedData() {
	// device data
	MariaDbClient.Exec("insert into device(id, name) values (?, ?)", 1, "mobile")
	MariaDbClient.Exec("insert into device(id, name) values (?, ?)", 2, "tablet")
	MariaDbClient.Exec("insert into device(id, name) values (?, ?)", 3, "web")

	// country data
	MariaDbClient.Exec("insert into country(id, name, keyword) values (?, ?, ?)", 1, "India", "IN")
	MariaDbClient.Exec("insert into country(id, name, keyword) values (?, ?, ?)", 2, "United States of America", "US")

	// web requests
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 1, 1, 180)
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 1, 2, 80)
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 1, 3, 100)
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 2, 1, 150)
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 2, 2, 60)
	MariaDbClient.Exec("insert into web_request(country_id, device_id, count) values (?, ?, ?)", 2, 3, 60)
	// timespent
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 1, 1, 190)
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 1, 2, 40)
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 1, 3, 160)
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 2, 1, 175)
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 2, 2, 50)
	MariaDbClient.Exec("insert into time_spent(country_id, device_id, count) values (?, ?, ?)", 2, 3, 40)
}
