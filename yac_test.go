package yac

import "testing"

func TestGetColorsByJPG(t *testing.T) {
	if err := GetColors("example.jpeg"); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log("All is okay")
}
