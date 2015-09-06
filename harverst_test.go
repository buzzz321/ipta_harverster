package main

import (
	"testing"
)

func TestGet_tokedata(t *testing.T) {
	res := get_tokendata("Hello=1", "=")

	if res != "1" {
		t.Errorf("get_tokendata returned %s extected 1")
	}
}
