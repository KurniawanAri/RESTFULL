package controllers

import (
	data "Training/Day15/selling/data"
	"encoding/json"
	"net/http"
)

func GetSelling(w http.ResponseWriter, r http.Request) {
	koneksi := Koneksi{}
	d := InitDB(koneksi.DB)
	defer d.Close()
	repo := &data.SellingData{d}
	selling := data.GetAll(repo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(selling)
	w.Write(mdl)
}
