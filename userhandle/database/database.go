package database

import (
	"net/http"
	"os"
	"regexp"
	"userhandle/communication"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	log "github.com/urishabh12/colored_log"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

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

func InsertUserRecord(new_user User) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}

	if new_user.Name == "" || new_user.Username == "" || new_user.Password == "" || new_user.Email == "" || new_user.Phone == "" {
		return "Please fill all the fields", http.StatusBadRequest, false
	}

	isEmailValid := govalidator.IsEmail(new_user.Email)
	if !isEmailValid {
		return "Invalid email", http.StatusBadRequest, false
	}

	isPhoneValid := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`).MatchString(new_user.Phone)
	if !isPhoneValid {
		return "Invalid phone number", http.StatusBadRequest, false
	}

	if new_user.UserType != "buyer" && new_user.UserType != "seller" {
		return "Invalid user type", http.StatusBadRequest, false
	}

	var existing_user User
	db.First(&existing_user, &User{Username: new_user.Username})

	//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
	if existing_user.Username == new_user.Username {
		return "This username is taken!", http.StatusConflict, false
	}
	if existing_user.Email == new_user.Email {
		return "This email is already registered!", http.StatusConflict, false
	}
	if existing_user.Phone == new_user.Phone {
		return "This phone number is already registered!", http.StatusConflict, false
	}

	new_user.Password = hashAndSalt([]byte(new_user.Password))

	result := db.Create(&new_user)
	if result.Error != nil || result.RowsAffected == 0 {
		return "Something went wrong on our side, please try again later", http.StatusInternalServerError, false
	} else {
		return "User created succesfully!", http.StatusCreated, true
	}
}

func CheckUserRecord(query_user communication.LoginRequest) (string, int, bool, *User) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}

	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})
	if existing_user.Username == query_user.Username && comparePasswords(existing_user.Password, []byte(query_user.Password)) {
		existing_user.Password = ""
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

	//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
	if existing_user.Username == query_user.Username {
		existing_user.Password = ""
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

	var existing_user User
	db.First(&existing_user, &User{Username: query_user_record.Username})
	if existing_user.Username != query_user_record.Username {
		log.Panic("[PANIC] JWT SECRET BREACH!!! Username:"+claims["username"].(string), "User record not found")
		return "Invalid malformed request", http.StatusForbidden, false
	}

	// check if password is present in the request, if not then we don't want to update it
	if new_info.Password == "" {
		new_info.Password = existing_user.Password
	} else {
		new_info.Password = hashAndSalt([]byte(new_info.Password))
	}

	if new_info.Address == "" {
		new_info.Address = existing_user.Address
	}

	updated_user_record := User{
		Username: new_info.Username,
		Password: new_info.Password,
		Address:  new_info.Address,
	}

	result := db.Model(&User{}).Where(&query_user_record).Updates(updated_user_record)
	if result.RowsAffected == 0 {
		log.Panic("[PANIC] JWT SECRET BREACH!!! Username:"+claims["username"].(string), result.Error)
		return "Invalid malformed request", http.StatusForbidden, false
	} else {
		return "User record updated!", http.StatusOK, true
	}
}

func DeleteUserRecord(query_user communication.DeleteRequest) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}

	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})
	if existing_user.Username == query_user.Username && comparePasswords(existing_user.Password, []byte(query_user.Password)) {
		result := db.Delete(&User{}, &User{Username: query_user.Username})
		if result.RowsAffected == 0 {
			return "Something went wrong on our side, please try again later", http.StatusInternalServerError, false
		} else {
			return "User record deleted succesfully!", http.StatusOK, true
		}
	} else {
		return "Invalid user records", http.StatusForbidden, false
	}
}
