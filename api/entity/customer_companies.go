package entity

type Customer_Companies struct {
	Company_Id   string      `json:"company_id" gorm:"primaryKey;AUTO_INCREMENT"`
	Company_Name string      `json:"company_name"`
	Customers    []Customers `json:"customers,omitempty" gorm:"foreignKey:CompanyId"`
}
