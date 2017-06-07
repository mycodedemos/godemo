package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:wxnacy@(192.168.99.100:32770)/ququ?charset=utf8mb4")
	checkErr(err)
	return db
}

func Query(sql string, args ...interface{}) {
	db := getDB()
	rows, err := db.Query(sql)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(args...)
		checkErr(err)
		fmt.Println(args[0])
	}
	db.Close()
}

func Execute(sql string, args ...interface{}) bool {
	db := getDB()
	stmt, err := db.Prepare(sql)
	checkErr(err)
	if err != nil {
		return false
	}
	res, err := stmt.Exec(args...)
	_ = res
	db.Close()
	checkErr(err)
	if err != nil {
		return false
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
