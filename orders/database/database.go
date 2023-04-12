package database

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

<<<<<<< HEAD
func GetDatabaseConnection() (*gorm.DB, error) {
	dsn := "storeuser:pass1234@tcp(127.0.0.1:3306)/cloudstore?charset=utf8mb4&parseTime=True&loc=Local"
	if db == nil { //If first time asking for database operations
=======
func InitDatabaseVaraiables() {

	envRequired := []string{"DATABASE_USERNAME", "DATABASE_PASSWORD", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME"}

	_, err := os.Stat(".env")
	if err == nil {
		secret, err := godotenv.Read()
		if err != nil {
			log.Panic("Error reading .env file")
		}

		for _, key := range envRequired {
			if secret[key] != "" {
				os.Setenv(key, secret[key])
			}
		}
	}

	for _, key := range envRequired {
		if os.Getenv(key) == "" {
			log.Panic("Environment variable " + key + " not set")
		}
	}
}

func GetDatabaseConnection() (*gorm.DB, error) {

	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	dsn := databaseUsername + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	if db == nil {
>>>>>>> b1491b0 (order mservice starting)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Error creating a connection to databse!", err)
			return nil, err
		}
<<<<<<< HEAD
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
=======
		db.AutoMigrate(&OrderKey{}) 
		db.AutoMigrate(&OrderItems{}) 
	}
	return db, nil
}
>>>>>>> b1491b0 (order mservice starting)
