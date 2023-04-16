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
	Order_id int32 `json:"order_id"`
}

type CreateOrderRequest struct {
	IDs []string `json:"ids"`
}

type UpdateOrderRequest struct {
}

type DeleteOrderRequest struct {
}

// type CartFetchRequest struct{
//    Product_ids  []string `json:"product_ids"`
// }
