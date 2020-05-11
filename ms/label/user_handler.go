package label

import (
	"fmt"
)

func userExist(uid int64) bool {
	db := dbConn()
	exist := false
	results, err := db.Query("SELECT idUser FROM User WHERE idUser=?", uid)
	if err != nil {
		exist = false
	} else {
		results.Next()
		var id int64
		err = results.Scan(&id)
		if err != nil {
			exist = false
		} else if id == uid {
			exist = true
			fmt.Println("user exist")
		}
	}
	defer results.Close()
	defer db.Close()
	return exist
}

func createUser(uid int64) bool {
	db := dbConn()
	created := false
	insForm, err := db.Prepare("INSERT INTO User(idUser) VALUES(?)")
	if err != nil {
		created = false
	}
	res, err := insForm.Exec(uid)
	if err != nil {
		created = false
	}
	fmt.Println(res)
	db.Close()
	return created
}

func deleteUserLabelDB(uid int64, idLabel int64) {
	db := dbConn()
	delete, err := db.Prepare("DELETE FROM Label_User WHERE User_idUser=? and Label_idLabel=?")
	if err != nil {
		fmt.Println("error ocurred")
		fmt.Println(err)
	}
	delete.Exec(uid,idLabel)
	defer delete.Close()
	defer db.Close()
}