package database

import (
	"net/http"
	"os"
	"products/communication"

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
		db.AutoMigrate(&Product{})
	}
	return db, nil
}

func GetAllProducts() (string, int, bool, []Product) {
	var query_result []Product
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	result := db.Find(&query_result)
	if result.Error != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	return "All products in store.", http.StatusOK, true, query_result
}

func GetAllSellerProducts(claims communication.LoginClaims) (string, int, bool, []Product) {
	query_prod := Product{
		SellerUsername: claims.Username,
	}
	var query_result []Product
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	result := db.Where(&query_prod).Find(&query_result)
	if result.Error != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	return "All products from this user", http.StatusOK, true, query_result
}

func GetProducts(ids []int) (string, int, bool, []Product) {
	var query_result []Product
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}

	for _, id := range ids {
		product := Product{}
		result := db.First(&product, id)
		if result.Error != nil {
			return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
		}
		query_result = append(query_result, product)
	}
	return "Products with given ids", http.StatusOK, true, query_result
}

func InsertProduct(product_details communication.CreateProductRequest, claims communication.LoginClaims) (string, int, bool, uint) {
	db_product_record := Product{
		Name:            product_details.Name,
		Limit:           product_details.Limit,
		SellerUsername:  claims.Username,
		Price:           product_details.Price,
		AvgRating:       0,
		NumberOfRatings: 0,
		Manufacturer:    product_details.Manufacturer,
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

func UpdateProduct(product_details communication.UpdateProductRequest, claims communication.LoginClaims) (string, int, bool, *Product) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	query_result := Product{}
	db.First(&query_result, product_details.ID)
	if query_result.ID != uint(product_details.ID) {
		return "Product with this product ID not found.", http.StatusForbidden, false, nil
	} else {
		if claims.Username != query_result.SellerUsername && claims.UserType != "admin" {
			log.Println("[WARN] Someone trying to break the system, wrong seller trying to edit non owned product.")
			return "You do not have permission to do this", http.StatusForbidden, false, nil
		}
		if product_details.Name != "" {
			query_result.Name = product_details.Name
		}
		if product_details.Price != 0 {
			query_result.Price = product_details.Price
		}
		if product_details.Limit != 0 {
			query_result.Limit = product_details.Limit
		}
		if product_details.Manufacturer != "" {
			query_result.Manufacturer = product_details.Manufacturer
		}
		result := db.Save(&query_result)
		if result.RowsAffected != 0 {
			return "Product updated succesfully", http.StatusOK, true, &query_result
		} else {
			return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
		}
	}
}

func DeleteProduct(product_details communication.DeleteProductRequest, claims communication.LoginClaims) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}
	query_result := Product{}
	db.First(&query_result, product_details.ID)
	if query_result.ID != uint(product_details.ID) {
		return "Product with this product ID not found.", http.StatusForbidden, false
	} else {
		if claims.Username != query_result.SellerUsername {
			log.Println("[WARN] Someone trying to break the system, wrong seller trying to edit non owned product.")
			return "You do not have permission to do this", http.StatusForbidden, false
		}
		query_result.ID = uint(product_details.ID)
		result := db.Delete(&query_result)
		if result.RowsAffected == 0 {
			return "Internal server error, please try again later", http.StatusInternalServerError, false
		} else {
			return "Product record deleted", http.StatusOK, true
		}
	}
}

func RateProduct(product_details communication.RateProductRequest, claims communication.LoginClaims) (string, int, bool, *Product) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	query_result := Product{}
	db.First(&query_result, product_details.ID)
	if query_result.ID != uint(product_details.ID) {
		return "Product with this product ID not found.", http.StatusForbidden, false, nil
	} else {
		if product_details.Rating < 1 || product_details.Rating > 5 {
			return "Rating must be between 1 and 5", http.StatusBadRequest, false, nil
		}
		query_result.NumberOfRatings++
		query_result.AvgRating = (query_result.AvgRating*float32(query_result.NumberOfRatings-1) + float32(product_details.Rating)) / float32(query_result.NumberOfRatings)
		result := db.Model(&query_result).Updates(query_result)
		if result.RowsAffected == 0 {
			return "Internal server error, please try again later", http.StatusInternalServerError, false, nil
		} else {
			return "Rated product succesfully", http.StatusOK, true, &query_result
		}
	}
}
