package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func create_table(cfg IPTAConfig) {
	fmt.Println("hello")
	fmt.Println(cfg.Main.Db_Sqlite_Filename)
	db, err := sql.Open("sqlite3", "./"+cfg.Main.Db_Sqlite_Filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	sqlStmt := fmt.Sprintf(`
	CREATE TABLE %s (
		id integer PRIMARY KEY AUTOINCREMENT,
		timestamp timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
		if_in varchar(10) DEFAULT NULL,
		if_out varchar(10) DEFAULT NULL,
		src_ip varchar(10) DEFAULT NULL,
		src_prt varchar(10) DEFAULT NULL,
		dst_ip varchar(10)  DEFAULT NULL,
		dst_prt varchar(10) DEFAULT NULL,
		proto varchar(10) DEFAULT NULL,
		action varchar(10) DEFAULT NULL,
		mac varchar(40) DEFAULT NULL);`, cfg.Main.Db_Table)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}
}

func delete_table(cfg IPTAConfig) {
	db, err := sql.Open("sqlite3", "./"+cfg.Main.Db_Sqlite_Filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
	sqlStmt := fmt.Sprintf(`
	DROP TABLE %s;`, cfg.Main.Db_Table)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}

}
