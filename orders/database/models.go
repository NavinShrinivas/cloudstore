package database

<<<<<<< HEAD
type Order_Item struct {
	Order_Id   string    `gorm:"primaryKey" json:"orderid"`
	Order      Order_Key `gorm:"foreignKey:Order_Id;references:Order_Id" json:"order"`
	Product_Id string    `json:"productid"`
	Rating     string    `json:"rating"`
}

type Order_Key struct {
	User_id  string `json:"userid"`
	Order_Id string `gorm:"primaryKey;autoIncrement:true" json:"orderid"`
=======
import (
	"gorm.io/gorm"
)

// -----Databse Models-----
type OrderItems struct {
	gorm.Model
	Order_id   string   `json:"order_id`
	Order_key  OrderKey `gorm:"foreignKey:Order_id;references:Order_id" json:"order_key"`
	Product_id string   `json:"productid"`
}

// As we are aiming a mico service arch, we dont need to strongly refrence every table and record, as long we can make sure the records wont be inserted that violates other services record we are fine
// That we are making sure by using auth and JWT
type OrderKey struct {
	Username string `json:"username"` //We dont need to refrence the actual User records
	Order_id string `json:"order_id gorm:"primaryKey;autoIncrement:true"`
>>>>>>> b1491b0 (order mservice starting)
}
