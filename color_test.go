// +build test

package base16

import (
	"github.com/shebang/base16"
	"testing"
)

func TestNewColor(t *testing.T) {
	var gotColor base16.Color
	var expectColor base16.Color

	// black
	gotColor = base16.NewColor("000000")
	expectColor = 0
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// white
	gotColor = base16.NewColor("ffffff")
	expectColor = 16777215
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
	// red
	gotColor = base16.NewColor("ff0000")
	expectColor = 16711680
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// green
	gotColor = base16.NewColor("00ff00")
	expectColor = 65280
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// blue
	gotColor = base16.NewColor("0000ff")
	expectColor = 255
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}

func TestNewColorErrorHandling(t *testing.T) {
	var gotColor base16.Color
	var expectColor base16.Color

	// unsupported returns no color
	gotColor = base16.NewColor("000000ff")
	expectColor = base16.NoColor
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}
