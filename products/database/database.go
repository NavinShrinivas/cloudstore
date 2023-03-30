package database

import (
	"net/http"
	"products/communication"

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

func SetProduct(product_details communication.CreateProductRequest, claims communication.LoginClaims) (string, int, bool, uint) {
	db_product_record := Product{
		Name:           product_details.Name,
		Limit:          product_details.Limit,
		Username:       claims.Username,
		Price:          product_details.Price,
		Avgrating:      0,
		Numberofrating: 0,
		Manufacturer:   product_details.Manufacturer,
	}
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, 0
	}
	result := db.Create(&db_product_record)
	if result.RowsAffected != 0 {
		return "Product inserted succesfully", http.StatusOK, true, db_product_record.ID
	} else {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, 0
	}
}

func EditProduct(product_details communication.EditProductRequest, claims communication.LoginClaims) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}
	query_result := Product{}
	db.First(&query_result, product_details.ID)
	if query_result.ID != uint(product_details.ID) {
		return "Product with this product ID not found.", http.StatusForbidden, false
	} else {
		if claims.Username != query_result.Username {
			return "You do not have permission to do this", http.StatusForbidden, false
			log.Println("[WARN] Someone trying to break the system, wrong seller trying to edit non owned product.")
		}
		query_result.Name = product_details.Name
		query_result.Price = product_details.Price
		query_result.Limit = product_details.Limit
		query_result.Manufacturer = product_details.Manufacturer
		result := db.Model(&query_result).Updates(query_result)
		if result.RowsAffected == 0 {
			return "Internal server error, please try again later", http.StatusInternalServerError, false
		} else {
			return "Updated product records succesfully", http.StatusOK, true
		}
	}
}

func GetAllProducts() (string,int,bool,[]Product){
   var query_result []Product 
   db,err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false,nil
	}
   result := db.Find(&query_result)
   if result.Error !=nil{
		return "Internal server error, please try again later.", http.StatusInternalServerError, false,nil
   }
   return "All products in store.",http.StatusOK,true,query_result
}


func GetAllSellerProducts(claims communication.LoginClaims) (string,int,bool,[]Product){
   query_prod := Product{
      Username : claims.Username,
   }
   var query_result []Product 
   db,err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false,nil
	}
   result := db.Where(&query_prod).Find(&query_result)
   if result.Error !=nil{
		return "Internal server error, please try again later.", http.StatusInternalServerError, false,nil
   }
   return "All products from this user",http.StatusOK,true,query_result
}
