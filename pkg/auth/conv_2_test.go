package auth

import "testing"

func TestGetPassword(t *testing.T) {
	err := GetPassword("漢字")
	if err != "漢字" {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
