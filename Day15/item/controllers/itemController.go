package controllers

import (
	data "Training/Day15/item/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	// Di dalam controller, buat file untuk koneksi
	// Setalah dari itemModel.go buat variabel koneksi
	koneksi := Koneksi{}
	d := InitDB(koneksi.DB)
	defer d.Close()
	repo := &data.ItemData{d}
	item := data.GetAll(repo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

//Setelah itemController.go pindah ke Folder routers buat itemRouter.go
