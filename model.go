package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Promotion struct {
	ID             string  `gorm:"primaryKey" json:"id"`
	Price          float32 `json:"price"`
	ExpirationDate string  `json:"expiration_date"`
}

func CreatePromotion(db *gorm.DB, p Promotion) error {
	return db.Create(&p).Error
}

func GetPromotionByID(db *gorm.DB, id string) (*Promotion, error) {
	var p Promotion
	if err := db.Where("id = ?", id).Find(&p).Error; err != nil {
		return nil, fmt.Errorf("error finding promotion by id. err: %w", err)
	}

	if p.ID == "" {
		return nil, gorm.ErrRecordNotFound
	}

	return &p, nil
}
