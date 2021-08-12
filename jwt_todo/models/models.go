package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

type User struct {
	gorm.Model
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type TokenDetails struct {
	gorm.Model
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	gorm.Model
	AccessUuid string
	UserId     uint64
}
