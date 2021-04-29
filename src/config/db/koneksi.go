package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Koneksi() *sql.DB {

	konfigDB := GetDBConf()

	//log.Fatal(connectionString)
	dbs, err := sql.Open("mysql", konfigDB)
	if err != nil {
		log.Fatal("ERR Connection")
	}
	return dbs
}
