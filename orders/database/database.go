package database

import (
	"net/http"
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
		db.AutoMigrate(&Order_Item{})
		db.AutoMigrate(&Order_Key{})
	}
	db.Set("gorm:auto_preload", true)
	return db, nil
}

func GetOrder(request communication.GetOrderRequest, claims communication.LoginClaims) (string, int, bool, *communication.Order) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	query_order := Order_Key{Order_Id: request.Order_id}
	order_info := Order_Key{}
	result := db.Preload("Order_Items").Find(&order_info, &query_order)
	if order_info.User_id != claims.Username {
		return "Invalid Request!", http.StatusBadRequest, false, nil
	}
	if result.Error != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	} else {
		order := communication.Order{Order_Id: order_info.Order_Id, User_id: order_info.User_id}
		for _, v := range order_info.Order_Items {
			details := communication.GetProductDetails(v.Product_Id)
			if details == nil {
				return "Invalid Request!", http.StatusBadRequest, false, nil
			} else {
				item := communication.Item{
					Product_Id: v.Product_Id,
					Order_Id:   v.Order_Id,
					Rating:     v.Rating,
					Details:    *details,
				}
				order.Items = append(order.Items, item)
			}
		}
		return "Found Order!", http.StatusOK, true, &order
	}
}

func GetAllOrders(claims communication.LoginClaims) (string, int, bool, []Order_Key) {

	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	if claims.UserType == "buyer" {
		// Get all orders of the buyer
		var all_order []Order_Key
		result := db.Preload("Order_Items").Find(&all_order, &Order_Key{User_id: claims.Username})
		if result.Error != nil {
			return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
		}
		return "Orders found!", http.StatusOK, true, all_order
	} else if claims.UserType == "seller" {
		//[TODO] this request will span to products micro service as well [DONE]
      //[TODO] This is by far the most unoptimised code block in this project!! please refactor if u ever can!
		var all_order []Order_Key
		var seller_order []Order_Key
		result := db.Preload("Order_Items").Find(&all_order)
		if result.Error != nil {
			return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
		}
		for _, v := range all_order {
			for _, vi := range v.Order_Items {
				details := communication.GetProductDetails(vi.Product_Id)
				if details != nil && details.Seller == claims.Username {
					seller_order = append(seller_order, v)
               //If one of the items is sellers we can move on, 
               //we dont want the same order to show up many times cus many of sellers products
               break
				} else if details == nil{
					return "Invalid Request!", http.StatusBadRequest, false, nil
				}
			}
		}
		return "Found all seller orders!", http.StatusOK, true, seller_order
	} else if claims.UserType == "admin" {
		var all_order []Order_Key
		result := db.Preload("Order_Items").Find(&all_order)
		if result.Error != nil {
			return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
		}
		return "Orders found!", http.StatusOK, true, all_order
	} else {
		// Invalid request
		return "Invalid Request!", http.StatusBadRequest, false, nil
	}
	return "Invalid Request!", http.StatusBadRequest, false, nil
}

func InsertOrder(orderDetails communication.CreateOrderRequest, claims communication.LoginClaims) (string, int, bool, *Order_Key) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	insert_record := Order_Key{User_id: claims.Username}
	result := db.Create(&insert_record)
	if result.Error != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	return "Intialized an order record!", http.StatusOK, true, &insert_record

}

func InsertOrderItems(orderDetails communication.CreateOrderRequest, claims communication.LoginClaims, order_key *Order_Key) (string, int, bool, *Order_Key) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	if len(orderDetails.IDs) == 0 {
		return "Invalid request", http.StatusBadRequest, false, nil
	}
	for _, v := range orderDetails.IDs {
		order_item := Order_Item{
			Order_Id:   order_key.Order_Id,
			Product_Id: v,
			Rating:     0,
		}
		result := db.Create(order_item)
		if result.Error != nil {
			return "Invalid request", http.StatusBadRequest, false, nil
		}
	}
	order := Order_Key{}
	db.Preload("Order_Items").First(&order, &Order_Key{Order_Id: order_key.Order_Id})
	return "Created order!", http.StatusOK, true, &order
}

// func UpdateOrder(orderDetails communication.UpdateOrderRequest, claims communication.LoginClaims) (string, int, bool, Order_Key) {
// }
//
// func DeleteOrder(orderDetails communication.DeleteOrderRequest, claims communication.LoginClaims) (string, int, bool, Order_Key) {
// }
