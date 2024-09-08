package models

import "github.com/jinzhu/gorm"

type Package struct {
	ID               uint    `gorm:"primary_key" json:"id"`
	Name             string  `json:"name"`
	SenderName       string  `json:"sender"`
	ReceiverName     string  `json:"receiver"`
	SenderLocation   string  `json:"sender_location"`
	ReceiverLocation string  `json:"receiver_location"`
	Fee              float64 `json:"fee"`
	Weight           float64 `json:"weight"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Package{})
}
