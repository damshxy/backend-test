package models

type User struct {
    ID       int    `json:"id" gorm:"primary_key; not null"`
    Username string `json:"username" gorm:"type:varchar(255)"`
    Password string `json:"password" gorm:"type:varchar(255)"`
}