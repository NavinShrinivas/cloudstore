package database

type Order_Item struct {
	Order_Id   string    `gorm:"primaryKey" json:"orderid"`
	Order      Order_Key `gorm:"foreignKey:Order_Id;references:Order_Id" json:"order"`
	Product_Id string    `json:"productid"`
	Rating     string    `json:"rating"`
}

type Order_Key struct {
	User_id  string `json:"userid"`
	Order_Id string `gorm:"primaryKey;autoIncrement:true" json:"orderid"`
}
