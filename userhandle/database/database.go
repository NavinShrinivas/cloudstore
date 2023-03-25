package database

import (
	"net/http"
	"userhandle/communication"

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
		db.AutoMigrate(&User{})
	}
	return db, nil
}

func PutUserRecords(insert_user User) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}
	var existing_user User
	db.First(&existing_user, &User{Username: insert_user.Username})
	if existing_user.Username == insert_user.Username {
		//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
		return "This username is taken!", http.StatusForbidden, false
	}
	result := db.Create(&insert_user)
	if result.Error != nil || result.RowsAffected == 0 {
		return "Something went wrong on our side, please try again later", http.StatusInternalServerError, false
	} else {
		return "User created succesfully!", http.StatusOK, true
	}
}

func CheckUserRecords(query_user communication.LoginRequest) (string, int, bool, *User) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})
	if existing_user.Username == query_user.Username && existing_user.Password == query_user.Password {
		//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
		return "The user record is valid", http.StatusOK, true, &existing_user
	} else {
		return "Invalid user records", http.StatusForbidden, false, nil
	}
}

func UpdateUserRecord(new_info communication.EditRequest, claims map[string]interface{}) (string, int, bool) {
	//Here we can check if a person is trying to modify a non existant record, this means the JWT password HAS BEEN LEAKED!!!!
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}
	query_user_record := User{
		Username: claims["username"].(string),
	}
	updated_user_record := User{
		Username: new_info.Username,
		Password: new_info.Password,
	}
	result := db.Model(&User{}).Where(&query_user_record).Updates(updated_user_record)
	if result.RowsAffected == 0 {
		log.Panic("[PANIC] JWT SECRET LEAKED!!!!!!!!!!")
		return "Invalid malformed request", http.StatusForbidden, false
	} else {
		return "User record updated!", http.StatusOK, true
	}
}
