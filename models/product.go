package models

type Product struct {
    ID       int    `json:"id" gorm:"primary_key; not null"`
    Name     string `json:"name" gorm:"type:varchar(255)"`
    Price    int    `json:"price" gorm:"type:int"`
    Quantity int    `json:"quantity" gorm:"type:int"`
}