package model

type User struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Hash     string `json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
