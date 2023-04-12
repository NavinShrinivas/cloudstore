package database

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabaseConnection() (*gorm.DB, error) {
	dsn := "storeuser:pass1234@tcp(127.0.0.1:3306)/cloudstore?charset=utf8mb4&parseTime=True&loc=Local"
	if db == nil { //If first time asking for database operations
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Error creating a connection to databse!", err)
			return nil, err
		}
		db.AutoMigrate(&Product{})
	}
	return db, nil
}


func CreateOrder(, claims map[string]interface{}) string {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later."
	}
	var existing_order Order_Key
	db.First(&existing_order, &Order_Key{User_id: new_info.User_id})
	if existing_order.User_id == new_info.User_id {
		//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
		return "This order is already present"
	}
	result := db.Create(&new_info)
	if result.Error != nil || result.RowsAffected == 0 {
		return "Something went wrong on our side, please try again later"
	} else {
		return "Order created succesfully!"
	}

}
