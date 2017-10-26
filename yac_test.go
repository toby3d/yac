package yac

import "testing"

func TestTransform(t *testing.T) {
	if err := transform("example.jpeg"); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log("Done! See result in yac.jpg")
}
