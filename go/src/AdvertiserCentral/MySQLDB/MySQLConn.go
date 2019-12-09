package MySQLDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQLDB() *sql.DB{
	db, err :=
		sql.Open("mysql", "root:root1234*@tcp(127.0.0.1:3306)/abishek")
	if err != nil{
		panic(err.Error())
	}
	return db
}
