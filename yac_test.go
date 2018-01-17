package yac

import (
	"image"
	"os"
	"testing"
)

func TestAnalyzeFile(t *testing.T) {
	colors, err := Analyze("safedns.jpeg")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	t.Logf("Colors: %#v", colors)
}

func TestAnalyzeImage(t *testing.T) {
	fm, err := os.Open("safedns.jpeg")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer fm.Close()

	m, _, err := image.Decode(fm)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	colors, err := Analyze(m)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	t.Logf("Colors: %#v", colors)
}
