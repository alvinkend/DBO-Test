package mysqlmodels

import (
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	CustomerID  int    	  `json:"customer_id"`
	OrderDate 	time.Time `json:"order_date"`
	TotalAmount int    	  `json:"total_amount"`
}
