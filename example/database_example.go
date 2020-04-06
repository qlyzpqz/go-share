package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.Database

	db = sql.Open("mysql", "root:123456@/db?charset=utf8")

	stmt, ok := db.Prepare("select * from product where Fid=?")
	rows, ok := stmt.Query(123)
	for rows.Next() {
		rows.Scan()
	}
}
