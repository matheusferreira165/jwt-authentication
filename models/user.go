package models

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
