/*
 * PinArt Labels MS
 *
 * A labels microservice for PinArt system.
 *
 * API version: 1.0.0
 */

package swagger

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetAllLabels(db *sql.DB, w http.ResponseWriter, r *http.Request) []Label {
	labelList := make([]Label, 0)
	results, err := db.Query("SELECT idLabel FROM Label")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return labelList
	}
	for results.Next() {
		var id int64
		err = results.Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return labelList
		}
		labelList = append(labelList, GetLabelFromDB(db, int(id), w, r))
	}
	return labelList
}

func GetLabelFromDB(db *sql.DB, id int, w http.ResponseWriter, r *http.Request) Label {
	var lab Label
	// Label Properties
	results, err := db.Query("SELECT name, description FROM Label WHERE idLabel=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return lab
	}
	results.Next()
	err = results.Scan(&lab.Name, &lab.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return lab
	}
	// Related Labels
	results, err = db.Query("SELECT Label_id1 as id from Label_relation where Label_idLabel =? union select Label_idLabel as id from Label_relation where Label_id1 = ?", id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return lab
	}

	list := make([]int64, 1)
	for results.Next() {
		var val int64
		err = results.Scan(&val)
		list = append(list, val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return lab
		}
	}
	lab.Id = int64(id)
	lab.RelatedLabels = list
	return lab
}

func updateLabelRelations(label Label, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	currentLabel := GetLabelFromDB(db, int(label.Id), w, r)
	forUnlink := Difference(currentLabel.RelatedLabels, label.RelatedLabels)
	forLink := Difference(label.RelatedLabels, currentLabel.RelatedLabels)
	for _, labelId := range forLink {
		linkLabel(label.Id, labelId, db, w)
	}
	for _, labelId := range forUnlink {
		unlinkLabel(label.Id, labelId, db, w)
	}
}

func linkLabel(id1 int64, id2 int64, db *sql.DB, w http.ResponseWriter) {
	insForm, err := db.Prepare("INSERT INTO Label_relation(Label_id1, label_idLabel) VALUES(?,?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := insForm.Exec(id1, id2)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(res)
}

func unlinkLabel(id1 int64, id2 int64, db *sql.DB, w http.ResponseWriter) {
	insForm, err := db.Prepare("DELETE FROM Label_relation WHERE (Label_id1 = ? AND label_idLabel = ?) OR (Label_id1 = ? AND label_idLabel = ?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := insForm.Exec(id1, id2, id2, id1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(res)
}

// Set Difference: A - B
func Difference(a, b []int64) (diff []int64) {
	m := make(map[int64]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
