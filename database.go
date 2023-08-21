package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(host, user, pass, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // use default config
	if err != nil {
		return nil, fmt.Errorf("error openning connetion to DB: %w", err)
	}

	return db, nil
}

func clearTable(db *gorm.DB) error {
	var recordCount int64
	if err := db.Model(&Promotion{}).Count(&recordCount).Error; err != nil {
		return fmt.Errorf("error counting records: %w", err)
	}

	if recordCount > 0 {
		if err := db.Migrator().DropTable(&Promotion{}); err != nil {
			return fmt.Errorf("error dropping table: %w", err)
		}

		if err := db.AutoMigrate(&Promotion{}); err != nil {
			return fmt.Errorf("error automigrating: %w", err)
		}
	}

	return nil
}
