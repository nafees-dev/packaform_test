package entity

type Order_Items struct {
	Id           string       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId      string       `json:"order_id" gorm:"foreignKey:Order;references:Id"`
	PricePerUnit string       `json:"price_per_unit"`
	Quantity     string       `json:"quantity"`
	Product      string       `json:"product"`
	Deliveries   []Deliveries `json:"Deliveries,omitempty" gorm:"foreignKey:OrderItemId"`
}
