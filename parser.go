package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"gorm.io/gorm"
)

const workerCount = 10

func parseCaseStudyCSV(db *gorm.DB, filename string) error {
	// before parsing clear database table.
	clearTable(db)

	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	var wg sync.WaitGroup
	jobs := make(chan []string, workerCount)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(db, &wg, jobs)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err != nil {
				close(jobs)
				break
			}
			jobs <- record
		}
	}()

	wg.Wait()
	log.Printf("File parsed successfully. Time took: %s \n", time.Since(start).String())
	return nil
}

func worker(db *gorm.DB, wg *sync.WaitGroup, jobs <-chan []string) {
	defer wg.Done()
	for record := range jobs {
		id := record[0]
		amount, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			continue
		}
		date := record[2]

		prm := Promotion{
			ID:             id,
			Price:          float32(amount),
			ExpirationDate: date,
		}
		if err := CreatePromotion(db, prm); err != nil {
			log.Printf("Error creating promotion in DB: %v", err.Error())
		}
	}
}
