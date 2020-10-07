package gobase16

import (
	// "github.com/shebang/gobase16"
	"testing"
)

func TestNewColor(t *testing.T) {
	var gotColor Color
	var expectColor Color

	// black
	gotColor = NewColor("000000")
	expectColor = 0
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// white
	gotColor = NewColor("ffffff")
	expectColor = 16777215
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
	// red
	gotColor = NewColor("ff0000")
	expectColor = 16711680
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// green
	gotColor = NewColor("00ff00")
	expectColor = 65280
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// blue
	gotColor = NewColor("0000ff")
	expectColor = 255
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}

func TestNewColorErrorHandling(t *testing.T) {
	var gotColor Color
	var expectColor Color

	// unsupported returns no color
	gotColor = NewColor("000000ff")
	expectColor = NoColor
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}
