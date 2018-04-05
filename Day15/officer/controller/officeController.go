package controller

import (
	data "Training/Day15/officer/data"
	"encoding/json"
	"net/http"
)

func GetOffice(w http.ResponseWriter, r *http.Request) {
	koneksi := Koneksi{}
	d := InitDB(koneksi.DB)
	defer d.Close()
	repo := &data.OfficeData{d}
	officer := data.GetAll(repo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)
	w.Write(mdl)
}
