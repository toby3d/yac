package yac

import "testing"

func TestAnalyzeFile(t *testing.T) {
	colors, err := Analyze("safedns.jpeg")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	t.Logf("Colors: %#v", colors)
}
