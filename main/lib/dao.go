package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUsername := "root"
	dbName := "tokens"
	db, err := sql.Open(dbDriver, dbUsername+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}


type Token struct {
	id int
	token string
}

func generateString() string {
	return "1234567"
}

//create new token and save to DB
func CreateToken() string {
	newToken := generateString()
	db := DbConn()
	insertToken, err := db.Prepare("INSERT INTO token (token, createdDate) values (?, now())")
	if err != nil {
		panic(err.Error())
	}
	insertToken.Exec(newToken)
	defer db.Close()
	return newToken
}

//Search for the Token to see if it exists
func TokenExists(searchToken string) bool {
	db := DbConn()
	selectString, err := db.Query("SELECT count(*) FROM token WHERE token=?", searchToken)
	if err != nil {
		panic(err.Error())
	}
	var count int

	for selectString.Next() {
		err = selectString.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	foundToken := false
	if count > 0 {
		foundToken = true
	}
 	return foundToken
}

func PathCount(path string) int {
	db := DbConn()
	//insert into used DB
	insertPath, err := db.Prepare("INSERT INTO used (path, usedDate) values (?, now())")
	if err != nil {
		panic(err.Error())
	}
	insertPath.Exec(path)

	//count how many total rows have been written
	selectString, err := db.Query("SELECT count(*) FROM used WHERE path=?", path)
	if err != nil {
		panic(err.Error())
	}
	var count int

	for selectString.Next() {
		err = selectString.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	return count

}