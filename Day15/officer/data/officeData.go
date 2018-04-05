package data

import (
	mdl "Training/Day15/officer/models"
	"database/sql"
)

type OfficeData struct {
	DB *sql.DB
}

func GetAll(db *OfficeData) []mdl.Office {
	result, err := db.DB.Query("SELECT OfficerCode,OfficerName from tblOfficer")
	if err != nil {
		return nil
	}
	defer result.Close()
	office := []mdl.Office{}
	for result.Next() {
		var o mdl.Office
		if err := result.Scan(&o.OfficerCode, &o.OfficerName); err != nil {
			return nil
		}
		office = append(office, o)
	}
	return office
}
