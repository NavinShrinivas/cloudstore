package communication

import (
	"github.com/golang-jwt/jwt"
)

// -----Communication models-----

type LoginClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserType string `json:"usertype"`
	// type : buyer|seller|admin
	jwt.StandardClaims
}

type AuthResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Claims  LoginClaims `json:"claims"`
}

type GetOrderRequest struct {
	Order_id int `json:"order_id"`
}

type CreateOrderRequest struct {
	IDs []int `json:"ids"`
}

type UpdateOrderRequest struct {
}

type DeleteOrderRequest struct {
}

// type CartFetchRequest struct{
//    Product_ids  []string `json:"product_ids"`
// }

// -----IPC Communication structs----
// ITZ MESSY DO NOT TOUCH IT
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Limit        int     `json:"limit"`
	Manufacturer string  `json:"manufacturer"`
	Seller       string  `json:"seller_username"`
}

type Products struct {
	Status   bool      `json:"status"`
	Message  string    `json:"message"`
	Products []Product `json:"products"`
}

type FetchProductsRequest struct {
	IDs []int `json:"ids"`
}

type Item struct {
	Order_Id   int     `json:"orderid"`
	Product_Id int     `json:"productid"`
	Rating     int32   `json:"rating"`
	Details    Product `json:"details"`
}

type Order struct {
	User_id  string `json:"userid"`
	Order_Id int    `json:"orderid"`
	Items    []Item `json:"items"`
}
