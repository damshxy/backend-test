package models

type Order struct {
    ID         int    `json:"id" gorm:"primary_key; not null"`
    Quantity   int    `json:"quantity" gorm:"type:int"`
    TotalPrice int    `json:"total_price" gorm:"type:int"`
    Status     string `json:"status" gorm:"type:varchar(255)"`
}