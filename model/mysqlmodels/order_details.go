package mysqlmodels

type OrderDetail struct {
	ID          int       `json:"id"`
	OrderID  	int    	  `json:"order_id"`
	ProductID 	int 	  `json:"product_id"`
	Qty 		int    	  `json:"qty"`
	OrderPrice 	int 	  `json:"order_price"`
}