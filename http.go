package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (a *App) GetPromotionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	prm, err := GetPromotionByID(a.db, id)
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("promotion not found"))
		return
	}

	promotionJson, err := json.Marshal(&prm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marhaling"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(promotionJson)
}
