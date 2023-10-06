package request

import "time"

type CustomerReq struct {
	ID              int     `json:"id" uri:"id"`
	CustomerName    string  `json:"customer_name"`
	Email           string  `json:"email"`
	ShippingAddres 	string  `json:"shipping_addres"`
	Password        string  `json:"password"`
	Page            int     `form:"page" json:"-"`
	Size            int     `form:"size" json:"-"`
}

type CustomerDetailReq struct {
	ID          int        `json:"id" uri:"id"`
	CustomerID  int    	   `json:"customer_id"`
	Gender 		string 	   `json:"gender"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	Page        int        `form:"page" json:"-"`
	Size        int        `form:"size" json:"-"`
}

type OrderReq struct {
	ID          int       `json:"id" uri:"id"`
	CustomerID  int    	  `json:"customer_id"`
	OrderDate 	time.Time `json:"order_date"`
	TotalAmount int    	  `json:"total_amount"`
	Page        int       `form:"page" json:"-"`
	Size        int       `form:"size" json:"-"`
}

type OrderDetailReq struct {
	ID          int       `json:"id" uri:"id"`
	OrderID  	int    	  `json:"order_id"`
	ProductID 	int 	  `json:"product_id"`
	Qty 		int    	  `json:"qty"`
	OrderPrice 	int 	  `json:"order_price"`
	Page        int       `form:"page" json:"-"`
	Size        int       `form:"size" json:"-"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenStructure struct {
	CustomerID int    `json:"customer_id"`
	Email  	   string `json:"email"`
	Role   	   string `json:"role"`
}