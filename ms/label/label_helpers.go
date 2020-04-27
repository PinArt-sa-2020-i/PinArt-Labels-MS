/*
 * PinArt Labels MS
 *
 * A labels microservice for PinArt system.
 *
 * API version: 1.0.0
 */

package label

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "labelms"
	dbPass := "2020i"
	dbName := "tcp(pinart-labels-db:3306)/labels" // "tcp(127.0.0.1:3306)/labels" / //
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	if err != nil {
		log.Panic(err.Error())
		panic(err.Error())
	}
	return db
}

func getBoardRelatedLabels(idBoard int64, db *sql.DB, w http.ResponseWriter, r *http.Request) []Label {
	labelList := make([]Label, 0)
	// gets the related labels id
	results, err := db.Query("SELECT Label_id FROM Board_Label WHERE Board_id = ?", idBoard)
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
		labelList = append(labelList, GetLabelFromDB(db, int(id), w))
	}

	return labelList
}

func getUserRelatedLabels(idUser int64, db *sql.DB, w http.ResponseWriter, r *http.Request) []Label {
	labelList := make([]Label, 0)
	// gets the related labels id
	results, err := db.Query("SELECT Label_idLabel FROM Label_User WHERE User_idUser = ?", idUser)
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
		labelList = append(labelList, GetLabelFromDB(db, int(id), w))
	}

	return labelList
}

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
		labelList = append(labelList, GetLabelFromDB(db, int(id), w))
	}
	return labelList
}

func GetLabelFromDB(db *sql.DB, id int, w http.ResponseWriter) Label {
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

func getLabels(ids []int, db *sql.DB, w http.ResponseWriter) []Label {
	list := make([]Label, 0)
	for _, id := range ids {
		list = append(list, GetLabelFromDB(db, int(id), w))
	}
	return list
}

func updateLabelRelations(label Label, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	currentLabel := GetLabelFromDB(db, int(label.Id), w)
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

func linkBoard(idBoard int64, idLabels []int, db *sql.DB, w http.ResponseWriter) {
	for _, label := range idLabels {
		insForm, err := db.Prepare("INSERT INTO Board_Label(Label_id, Board_id) VALUES(?,?)")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		res, err := insForm.Exec(label, idBoard)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(res)
	}
}

func linkUser(idUser int64, idLabels []int, db *sql.DB, w http.ResponseWriter) {
	for _, label := range idLabels {
		insForm, err := db.Prepare("INSERT INTO Label_User(User_idUser, Label_idLabel) VALUES(?,?)")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		res, err := insForm.Exec(idUser, label)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(res)
	}
}

func unlinkBoard(idBoard int64, idLabel int64, db *sql.DB, w http.ResponseWriter) {
	insForm, err := db.Prepare("DELETE FROM Board_Label WHERE (Label_id = ? AND Board_id = ?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := insForm.Exec(idLabel, idBoard)
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
