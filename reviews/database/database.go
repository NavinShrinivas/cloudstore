package database

import (
	"os"
	"reviews/communication"

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
		db.AutoMigrate(&Review{})
	}
	return db, nil
}

func GetReview(review_id string, claims communication.LoginClaims) (string, int, bool, Review_Key) {
	// check if the review belongs to the user (ie the user is the buyer or the seller) or the user is an admin
}

func GetAllReviews(claims communication.LoginClaims) (string, int, bool, []Review_Key) {
	if claims.UserType == "buyer" {
		// Get all reviews of the buyer
	} else if claims.UserType == "seller" {
		// Get all reviews of the seller
	} else if claims.UserType == "admin" {
		// Get all reviews of all users
	} else {
		// Invalid request
	}
}

func InsertReview(reviewDetails communication.CreateReviewRequest, claims communication.LoginClaims) (string, int, bool, Review_Key) {
}

func UpdateReview(reviewDetails communication.UpdateReviewRequest, claims communication.LoginClaims) (string, int, bool, Review_Key) {
}

func DeleteReview(reviewDetails communication.DeleteReviewRequest, claims communication.LoginClaims) (string, int, bool, Review_Key) {
}
