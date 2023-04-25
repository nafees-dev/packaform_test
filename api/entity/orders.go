package entity

type Order struct {
	Id         string    `json:"id"`
	CreatedAt  string    `json:"created_at"`
	OrderName  string    `json:"order_name"`
	CustomerId string    `json:"customer_id"`
	Customers  Customers `json:"Customers,omitempty" gorm:"foreignKey:CustomerId;references:UserId"`
	OrderItems []Order_Items `json:"OrderItems,omitempty" gorm:"foreignKey:OrderId"`
}
