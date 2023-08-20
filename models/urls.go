package models

import "time"

type Urls struct {
	Id         uint      `json:"id" grom:"primaryKey"`
	Url        string    `json:"url" gorm:"not null" validate:"required,min=3"`
	SlugUrl    string    `json:"slugUrl" gorm:"not null"`
	ExpireDate time.Time `json:"expireDate" gorm:"not null"`
	IsExpired  bool      `json:"isExpired" gorm:"not null;default:false"`
	CreatedAt  time.Time `json:"createdAt"`
}
