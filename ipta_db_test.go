package main

import (
	"testing"
)

/*
	Db_Type            string
	Db_Name            string
	Db_Host            string
	Db_User            string
	Db_Pass            string
	Db_Table           string
	Db_Sqlite_Filename string
*/

func make_cfg() *IPTAConfig {
	return &IPTAConfig{Main: Main{
		Db_Type:            "mysql",
		Db_Name:            "IPTA",
		Db_Host:            "localhost",
		Db_User:            "user",
		Db_Pass:            "******",
		Db_Table:           "IPTA",
		Db_Sqlite_Filename: "",
	}}
}

/*
	action string
	inIf   string
	outIf  string
	mac    string
	src    string
	dst    string
	proto  string
	spt    string
	dpt    string
	date   time.Time
*/

func TestAdd_items(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cfg := make_cfg()
	var items [2]FWItem

	items[0] = FWItem{action: "", inIf: "", outIf: "", mac: ""}
	add_item(*cfg, db, items[0])

	//	if res != "1" {
	//		t.Errorf("get_tokendata returned %s extected 1")
	//	}
}
