package entity

type Customers struct {
	UserId             string             `json:"user_id" gorm:"primaryKey;AUTO_INCREMENT"`
	Login              string             `json:"login"`
	Password           string             `json:"-"`
	Name               string             `json:"name"`
	CompanyId          string             `json:"company_id"`
	CreditCards        string             `json:"-"`
	Customer_Companies Customer_Companies `json:"Customer_Companies,omitempty" gorm:"foreignKey:CompanyId"`
}
