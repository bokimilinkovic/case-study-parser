package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const intervalFreq = 30
const filename = "promotions.csv"

type App struct {
	db *gorm.DB
}

func main() {
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASS", "postgres")
	name := getenv("DB_NAME", "casestudy")
	port := getenv("DB_PORT", "5433")

	db, err := ConnectToDB(host, user, pass, name, port)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{db: db}

	if err := db.AutoMigrate(&Promotion{}); err != nil {
		log.Fatalf("error automigrating table: %v", err.Error())
	}

	parseCaseStudyCSV(db, filename)

	ticker := time.NewTicker(intervalFreq * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				parseCaseStudyCSV(db, filename)
			}
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/promotions/{id}", app.GetPromotionByID)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getenv(envname, dflt string) string {
	env := os.Getenv(envname)
	if env == "" {
		return dflt
	}

	return env
}
