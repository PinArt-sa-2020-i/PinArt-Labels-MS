package label

import (
	"fmt"
	"net/http"
)

func labelSearch(fragment string, w http.ResponseWriter) []Label {
	db := dbConn()
	labelList := make([]Label, 0)
	fmt.Println(fragment)
	fragment = "%" + fragment + "%"
	fmt.Println(fragment)
	results, err := db.Query("SELECT idLabel FROM Label WHERE name like ? or description like ?", fragment, fragment)
	if err != nil {
		http.Error(w, "No se encuentra la etiqueta", http.StatusInternalServerError)
		return labelList
	}
	for results.Next() {
		var id int64
		err = results.Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return labelList
		}
		labelList = append(labelList, GetLabelFromDB(db, int(id), w))
	}
	return labelList
}
