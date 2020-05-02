/*
 * PinArt Labels MS
 *
 * A labels microservice for PinArt system.
 *
 * API version: 1.0.0
 */

package label

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func AddLabel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db := dbConn()
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Printf("%s", body)
	var theLabel Label
	err = json.Unmarshal(body, &theLabel)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	insForm, err := db.Prepare("INSERT INTO Label(name, description) VALUES(?,?)")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := insForm.Exec(theLabel.Name, theLabel.Description)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(result)
	q, err := db.Query("SELECT LAST_INSERT_ID()")
	var id Id
	q.Next()
	q.Scan(&id.Id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// write all the related labels
	for _, relatedID := range theLabel.RelatedLabels {
		linkLabel(id.Id, relatedID, db, w)
	}
	theLabel.Id = id.Id
	js, err := json.Marshal(theLabel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func DeleteLabel(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	idlabel, val := getCodeLabel(r, 0)
	fmt.Println(val)
	if idlabel == 0 {
		log.Printf("Error reading param: %v", idlabel)
		http.Error(w, "can't read params", http.StatusBadRequest)
		return
	}
	var theLabel Label
	theLabel.Id = int64(idlabel)

	delete, err := db.Prepare("DELETE FROM Label WHERE idLabel=?")
	if err != nil {
		http.Error(w, "can't delete label", http.StatusInternalServerError)
	}
	delete.Exec(theLabel.Id)
	js, err := json.Marshal(theLabel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	w.Write(js)
}

func GetLabel(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	var js []byte
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		var labelList []Label
		labelList = GetAllLabels(db, w, r)
		fmt.Println(labelList)
		js, err = json.Marshal(labelList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// gets one label from db
		label := GetLabelFromDB(db, id, w)
		js, err = json.Marshal(label)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func UpdateLabel(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Printf("%s", body)
	var theLabel Label
	err = json.Unmarshal(body, &theLabel)

	update, err := db.Prepare("UPDATE Label SET name=?, description=? WHERE idLabel=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(theLabel.Name, theLabel.Description, theLabel.Id)
	updateLabelRelations(theLabel, db, w, r)
	js, err := json.Marshal(theLabel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	w.Write(js)
}

func SearchLabel(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var labelList []Label
	// get the fragment
	keys, ok := r.URL.Query()["fragment"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "No fragment value given", http.StatusInternalServerError)
		return
	}
	fragment := keys[0]
	// calls handler
	labelList = labelSearch(fragment, w)
	response, err := json.Marshal(labelList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
