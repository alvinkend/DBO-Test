package mysqlmodels

type Customer struct {
	ID              int       `json:"id"`
	CustomerName    string    `json:"customer_name"`
	Email           string    `json:"email"`
	ShippingAddres 	string    `json:"shipping_addres"`
	Password        string `json:"password"`
}