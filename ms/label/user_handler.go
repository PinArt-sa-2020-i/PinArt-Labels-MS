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
	return created
}
