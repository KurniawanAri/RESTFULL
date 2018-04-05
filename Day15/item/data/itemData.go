package data

import (
	mdl "Training/Day15/item/models"
	"database/sql"
)

type ItemData struct {
	DB *sql.DB
}

func GetAll(db *ItemData) []mdl.Item {
	//fmt.Println(db.DB)
	result, err := db.DB.Query("SELECT ItemName from tblItem")
	if err != nil {
		return nil
	}
	defer result.Close()
	//fmt.Println(result)
	item := []mdl.Item{}
	for result.Next() {
		var i mdl.Item
		if err := result.Scan(&i.ItemName); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}

//Setelah selesai pindah ke Folder controllers file itemController.go
