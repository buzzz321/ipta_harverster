package main

import (
	"testing"
)

var testdata = "May  4 06:35:08 zathras kernel: [1749207.946614] IPT: CBLK IN=eth0 OUT= MAC=02:00:00:00:00:00:00:00:00:21:29:40:08:00 SRC=11.22.222.111 DST=111.222.22.222 LEN=40 TOS=0x00 PREC=0x00 TTL=56 ID=10028 PROTO=TCP SPT=55844 DPT=445 WINDOW=512 RES=0x00 SYN URGP=0"

func TestGet_tokedata(t *testing.T) {
	res := get_tokendata("Hello=1", "=")

	if res != "1" {
		t.Errorf("get_tokendata returned %s extected 1")
	}
}

func TestParse_line(t *testing.T) {
	valid, res := parse_line(testdata)

	if !valid {
		t.Errorf("parse_line returned %t extected true", valid)
	}

	if res.action != "CBLK" {
		t.Errorf("parse_line returned %s extected CBLK", res.action)
	}
}
