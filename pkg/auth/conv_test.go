package auth

import "testing"

func TestAGetID(t *testing.T) {
	err := GetID("漢字")
	if err != "漢字" {
		t.Log("error should be nil", err)
		t.Fail()
	}
}
