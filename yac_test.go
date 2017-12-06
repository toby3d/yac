package yac

import "testing"

func TestTransform(t *testing.T) {
	colors, err := Analyze("example.jpeg")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	for i := range colors {
		r, g, b, _ := colors[i].RGBA()
		t.Log("R:", (r >> 8))
		t.Log("G:", (g >> 8))
		t.Log("B:", (b >> 8))
		t.Log("======")
	}
}
