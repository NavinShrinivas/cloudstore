package database

import (
	"orders/communication"
	"os"

	log "github.com/urishabh12/colored_log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabaseConnection() (*gorm.DB, error) {

	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	dsn := databaseUsername + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	if db == nil {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Error creating a connection to databse!", err)
			return nil, err
		}
		db.AutoMigrate(&Order_Key{}, &Order_Item{})
	}
	return db, nil
}

func GetOrder(order_id string, claims communication.LoginClaims) (string, int, bool, Order_Key) {
	// check if the order belongs to the user (ie the user is the buyer or the seller) or the user is an admin

}

func GetAllOrders(claims communication.LoginClaims) (string, int, bool, []Order_Key) {
	if claims.UserType == "buyer" {
		// Get all orders of the buyer
	} else if claims.UserType == "seller" {
		// Get all orders of the seller
	} else if claims.UserType == "admin" {
		// Get all orders of all users
	} else {
		// Invalid request
	}
}

func InsertOrder(orderDetails communication.CreateOrderRequest, claims communication.LoginClaims) (string, int, bool, Order_Key) {
}

func UpdateOrder(orderDetails communication.UpdateOrderRequest, claims communication.LoginClaims) (string, int, bool, Order_Key) {
}

func DeleteOrder(orderDetails communication.DeleteOrderRequest, claims communication.LoginClaims) (string, int, bool, Order_Key) {
}
