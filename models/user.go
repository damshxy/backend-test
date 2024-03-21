package models

type User struct {
    ID       int    `json:"id" gorm:"primary_key; not null"`
    Username string `json:"username" gorm:"type:varchar(255)" validate:"required"`
    Password string `json:"password" gorm:"type:varchar(255)" validate:"required"`
}

type ResetPassword struct {
	Username string `json:"username" validate:"required, username"`
	Token string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}