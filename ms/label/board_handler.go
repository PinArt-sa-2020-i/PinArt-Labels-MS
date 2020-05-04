package label

import 	"fmt"

func boardExist(bid int64) bool {
	db := dbConn()
	exist := false
	results, err := db.Query("SELECT idBoard FROM Board WHERE idBoard=?", bid)
	if err != nil {
		exist = false
	} else {
		results.Next()
		var id int64
		err = results.Scan(&id)
		if err != nil {
			exist = false
		} else if id == bid {
			exist = true
		}
	}
	defer results.Close()
	defer db.Close()
	return exist
}

func createBoard(bid int64) bool {
	db := dbConn()
	created := false
	insForm, err := db.Prepare("INSERT INTO Board(idBoard) VALUES(?)")
	if err != nil {
		created = false
	}
	res, err := insForm.Exec(bid)
	if err != nil {
		created = true
	}
	fmt.Println(res)
	defer insForm.Close()
	defer db.Close()
	return created
}
