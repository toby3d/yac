package yac

import (
	"testing"
)

func TestTransform(t *testing.T) {
	colors, err := transform("example.jpeg")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	for i := range colors {
		t.Log("R:", colors[i].R)
		t.Log("G:", colors[i].G)
		t.Log("B:", colors[i].B)
		t.Log("A:", colors[i].A)
		t.Log("======")
	}
}
