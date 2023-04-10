package database

import (
	"net/http"
	"os"
	"userhandle/communication"

	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabaseVaraiables() {

	secret, err := godotenv.Read()
	if err != nil {
		log.Panic("Error reading .env file")
	}

	//Check if all the required environment variables are set
	if secret["DATABASE_USERNAME"] == "" {
		log.Panic("DATABASE_USERNAME not set in .env file")
	}
	if secret["DATABASE_PASSWORD"] == "" {
		log.Panic("DATABASE_PASSWORD not set in .env file")
	}
	if secret["DATABASE_HOST"] == "" {
		log.Panic("DATABASE_HOST not set in .env file")
	}
	if secret["DATABASE_PORT"] == "" {
		log.Panic("DATABASE_PORT not set in .env file")
	}
	if secret["DATABASE_NAME"] == "" {
		log.Panic("DATABASE_NAME not set in .env file")
	}

	os.Setenv("DATABASE_USERNAME", secret["DATABASE_USERNAME"])
	os.Setenv("DATABASE_PASSWORD", secret["DATABASE_PASSWORD"])
	os.Setenv("DATABASE_HOST", secret["DATABASE_HOST"])
	os.Setenv("DATABASE_PORT", secret["DATABASE_PORT"])
	os.Setenv("DATABASE_NAME", secret["DATABASE_NAME"])
}

func GetDatabaseConnection() (*gorm.DB, error) {

	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	log.Println("Connecting to database...")
	dsn := databaseUsername + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

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

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("Error in hashing password", err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("Error in comparing password", err)
		return false
	}
	return true
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
		return "This username is taken!", http.StatusConflict, false
	}
	insert_user.Password = hashAndSalt([]byte(insert_user.Password))
	result := db.Create(&insert_user)
	if result.Error != nil || result.RowsAffected == 0 {
		return "Something went wrong on our side, please try again later", http.StatusInternalServerError, false
	} else {
		return "User created succesfully!", http.StatusCreated, true
	}
}

func CheckUserRecords(query_user communication.LoginRequest) (string, int, bool, *User) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})
	if existing_user.Username == query_user.Username && comparePasswords(existing_user.Password, []byte(query_user.Password)) {
		return "The user record is valid", http.StatusOK, true, &existing_user
	} else {
		return "Invalid user records", http.StatusForbidden, false, nil
	}
}

func GetUserRecord(query_user User) (bool, *User) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return false, nil
	}
	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})
	if existing_user.Username == query_user.Username {
		//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
		return true, &existing_user
	} else {
		return false, nil
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
	new_info.Password = hashAndSalt([]byte(new_info.Password))
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
