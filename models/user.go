package model

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `gorm:"unique" json:"username" db:"username"`
	Email    string `gorm:"unique" json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"name"`
}
