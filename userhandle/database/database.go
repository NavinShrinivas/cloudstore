package database

import (
	"net/http"
	"os"
	"regexp"
	"userhandle/communication"

	"github.com/asaskevich/govalidator"
	log "github.com/urishabh12/colored_log"
	"golang.org/x/crypto/bcrypt"

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

	db.First(&existing_user, &User{Email: new_user.Email})
	if existing_user.Email == new_user.Email {
		return "This email is already registered!", http.StatusConflict, false
	}

	db.First(&existing_user, &User{Phone: new_user.Phone})
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

func GetUserRecord(query_user User) (bool, *UserInfo) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return false, nil
	}

	var existing_user User
	db.Select("name", "username", "email", "phone", "user_type", "address").First(&existing_user, &User{Username: query_user.Username})

	user_info := UserInfo{
		Name:     existing_user.Name,
		Username: existing_user.Username,
		Email:    existing_user.Email,
		Phone:    existing_user.Phone,
		UserType: existing_user.UserType,
		Address:  existing_user.Address,
	}

	//Need to do this check even though we have primary key as gorm add's it own primary key 'Id' making our entire primary key compostie and non uniuqe
	if existing_user.Username == query_user.Username {
		return true, &user_info
	} else {
		return false, nil
	}
}

func UpdateUserRecord(new_info communication.EditRequest, claims map[string]interface{}) (string, int, bool, *UserInfo) {

	//Here we can check if a person is trying to modify a non existant record, this means the JWT password HAS BEEN LEAKED!!!!
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false, nil
	}
	query_user_record := User{
		Username: claims["username"].(string),
	}

	var existing_user User
	db.First(&existing_user, &User{Username: query_user_record.Username})
	if existing_user.Username != query_user_record.Username {
		log.Panic("[PANIC] JWT SECRET BREACH!!! Username:"+claims["username"].(string), "User record not found")
		return "Invalid malformed request", http.StatusForbidden, false, nil
	}

	// check if password is correct
	if !comparePasswords(existing_user.Password, []byte(new_info.Password)) {
		log.Panic("[PANIC] JWT SECRET BREACH!!! Username:"+claims["username"].(string), "User record not found")
		return "Invalid malformed request", http.StatusForbidden, false, nil
	}

	// check the fields that are being updated
	if new_info.NewUsername == "" {
		new_info.NewUsername = existing_user.Username
	} else {
		var existing_user User
		db.First(&existing_user, &User{Username: new_info.NewUsername})
		if existing_user.Username == new_info.NewUsername {
			return "This username is taken!", http.StatusConflict, false, nil
		}
	}

	if new_info.NewEmail == "" {
		new_info.NewEmail = existing_user.Email
	} else {
		if !govalidator.IsEmail(new_info.NewEmail) {
			return "Invalid email", http.StatusBadRequest, false, nil
		}
		var existing_user User
		db.First(&existing_user, &User{Email: new_info.NewEmail})
		if existing_user.Email == new_info.NewEmail {
			return "This email is already registered!", http.StatusConflict, false, nil
		}
	}

	if new_info.NewPhone == "" {
		new_info.NewPhone = existing_user.Phone
	} else {
		isPhoneValid := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`).MatchString(new_info.NewPhone)
		if !isPhoneValid {
			return "Invalid phone number", http.StatusBadRequest, false, nil
		}
		var existing_user User
		db.First(&existing_user, &User{Phone: new_info.NewPhone})
		if existing_user.Phone == new_info.NewPhone {
			return "This phone number is already registered!", http.StatusConflict, false, nil
		}
	}

	// For now we are not allowing user type to be changed
	// if new_info.NewUserType == "" {
	// 	new_info.NewUserType = existing_user.UserType
	// } else {
	// 	if new_info.NewUserType != "buyer" && new_info.NewUserType != "seller" {
	// 		return "Invalid user type", http.StatusForbidden, false, nil
	// 	}
	// }

	if new_info.NewName == "" {
		new_info.NewName = existing_user.Name
	}

	if new_info.NewPassword != "" {
		new_info.NewPassword = hashAndSalt([]byte(new_info.NewPassword))
	} else {
		new_info.NewPassword = existing_user.Password
	}

	if new_info.NewAddress == "" {
		new_info.NewAddress = existing_user.Address
	}

	updated_user_record := User{
		Username: new_info.NewUsername,
		Password: new_info.NewPassword,
		Name:     new_info.NewName,
		Email:    new_info.NewEmail,
		Phone:    new_info.NewPhone,
		UserType: existing_user.UserType,
		Address:  new_info.NewAddress,
	}

	result := db.Model(&User{}).Where(&query_user_record).Updates(updated_user_record)
	if result.RowsAffected == 0 {
		log.Panic("[PANIC] JWT SECRET BREACH!!! Username:"+claims["username"].(string), result.Error)
		return "Invalid malformed request", http.StatusForbidden, false, nil
	} else {
		return "User record updated!", http.StatusOK, true, &UserInfo{
			Username: new_info.NewUsername,
			Name:     new_info.NewName,
			Email:    new_info.NewEmail,
			Phone:    new_info.NewPhone,
			UserType: existing_user.UserType,
			Address:  new_info.NewAddress,
		}
	}
}

func DeleteUserRecord(query_user communication.DeleteRequest) (string, int, bool) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return "Internal server error, please try again later.", http.StatusInternalServerError, false
	}

	var existing_user User
	db.First(&existing_user, &User{Username: query_user.Username})

	if existing_user.Username == "" {
		return "Invalid user records", http.StatusForbidden, false
	}

	if comparePasswords(existing_user.Password, []byte(query_user.Password)) {
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
