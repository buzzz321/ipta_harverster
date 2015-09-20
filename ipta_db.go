package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

func open_db(cfg IPTAConfig) *sql.DB {
	userinfo := fmt.Sprintf(`%s:%s@/%s`, cfg.Main.Db_User, cfg.Main.Db_Pass, cfg.Main.Db_Name)
	db, err := sql.Open(cfg.Main.Db_Type, userinfo)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return db
}

func create_db(cfg IPTAConfig) {
	db, err := sql.Open(cfg.Main.Db_Type, cfg.Main.Db_Sqlite_Filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

}

func create_table(cfg IPTAConfig) {

	db := open_db(cfg)
	defer db.Close()

	sqlStmt := fmt.Sprintf(`
	CREATE TABLE %s (
		id int(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
		timestamp timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
		if_in varchar(10) DEFAULT NULL,
		if_out varchar(10) DEFAULT NULL,
		src_ip int(10) unsigned DEFAULT NULL,
		src_prt int(10) unsigned DEFAULT NULL,
		dst_ip int(10) unsigned  DEFAULT NULL,
		dst_prt int(10) unsigned DEFAULT NULL,
		proto varchar(10) DEFAULT NULL,
		action varchar(10) DEFAULT NULL,
		mac varchar(40) DEFAULT NULL);`, cfg.Main.Db_Table)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}
}

func delete_table(cfg IPTAConfig) {
	db := open_db(cfg)
	defer db.Close()
	sqlStmt := fmt.Sprintf(`
	DROP TABLE %s;`, cfg.Main.Db_Table)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}
}

func clear_db(cfg IPTAConfig) {
	db := open_db(cfg)
	defer db.Close()

	sqlStmt := fmt.Sprintf(`DELETE FROM %s;	`, cfg.Main.Db_Table)

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		os.Exit(1)
	}

}

func add_item(cfg IPTAConfig, db *sql.DB, item FWItem) {

	sqlStmt := fmt.Sprintf(`
	INSERT INTO %s  ( timestamp, if_in, if_out, src_ip, src_prt, dst_ip, dst_prt, proto, action, mac) VALUES
( '%s', '%s', '%s', INET_ATON('%s'), '%s', INET_ATON('%s'),
'%s', '%s', '%s', '%s' )
	`, cfg.Main.Db_Table, item.date.Format("2006-01-02 15:04:05"), item.inIf, item.outIf, item.src, item.spt, item.dst, item.dpt, item.proto, item.action, item.mac)

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

func add_items(cfg IPTAConfig, db *sql.DB, items []FWItem) {

	sqlStmt := fmt.Sprintf(`
	INSERT INTO %s  (timestamp, if_in, if_out, src_ip, src_prt, dst_ip, dst_prt, proto, action, mac) VALUES
	`, cfg.Main.Db_Table)

	for item := range items {
		sqlStmt = strings.Join(sqlStmt, fmt.Sprintf(`
('%s', '%s', '%s', INET_ATON('%s'), '%s', INET_ATON('%s'),
'%s', '%s', '%s', '%s'),
	`, item.date.Format("2006-01-02 15:04:05"), item.inIf, item.outIf, item.src, item.spt, item.dst, item.dpt, item.proto, item.action, item.mac))
	}
	sqlStmt = strings.Join(sqlStmt[:-1], ";")

	fmt.Println(sqlStmt)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}
