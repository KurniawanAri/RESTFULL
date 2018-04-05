package data

import (
	mdl "Training/Day15/selling/models"
	"database/sql"
)

type SellingData struct {
	DB *sql.DB
}

func GetAll(db *SellingData) []mdl.Selling {
	result, err := db.DB.Query("SELECT Invoice from tblSelling")
	if err != nil {
		return nil
	}
	defer result.Close()
	selling := []mdl.Selling{}
	for result.Next() {
		var s mdl.Selling
		if err := result.Scan(&s.Invoice); err != nil {
			return nil
		}
		selling = append(selling, s)
	}
	return selling
}
