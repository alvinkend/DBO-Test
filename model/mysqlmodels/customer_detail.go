package mysqlmodels

import "time"

type CustomerDetail struct {
	ID          int        `json:"id"`
	CustomerID  int    	   `json:"customer_id"`
	Gender 		string 	   `json:"gender"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	OrderPrice 	int 	   `json:"order_price"`
}