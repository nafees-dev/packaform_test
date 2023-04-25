package entity

type Deliveries struct {
	Id                string `json:"id"`
	OrderItemId       string `json:"order_item_id" gorm:"foreignKey:OrderItemId;references:Id"`
	DeliveredQuantity string `json:"delivered_quantity"`
}
