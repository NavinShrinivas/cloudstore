package database

type Order_Item struct {
	Order_Id   int `json:"orderid"`
	Product_Id int    `json:"productid"`
	Rating     int32    `json:"rating"`
}

type Order_Key struct {
	User_id  string `json:"userid"`
	Order_Id int `gorm:"primaryKey;autoIncrement:true" json:"orderid"`
   Order_Items []Order_Item `gorm:"foreignKey:Order_Id;refrences:Order_Id" json:"items"`
}
