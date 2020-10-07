// +build test

package base16

import (
	"github.com/shebang/base16"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

// var base16TestData = `
// scheme: "Default Dark"
// author: "Chris Kempson (http://chriskempson.com)"
// base00: "181818"
// base01: "282828"
// base02: "383838"
// base03: "585858"
// base04: "b8b8b8"
// base05: "d8d8d8"
// base06: "e8e8e8"
// base07: "f8f8f8"
// base08: "ab4642"
// base09: "dc9656"
// base0A: "f7ca88"
// base0B: "a1b56c"
// base0C: "86c1b9"
// base0D: "7cafc2"
// base0E: "ba8baf"
// base0F: "a16946"
// `

func TestLoadBase16Scheme(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark.yaml")

	base16Scheme, err := base16.LoadBase16Scheme(testFile, ioutil.ReadFile)

	if err != nil {
		t.Errorf("expected no error err=%v", err)
	}

	gotString = base16Scheme.Author()
	expectString = "Chris Kempson (http://chriskempson.com)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = base16Scheme.Scheme()
	expectString = "Default Dark"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotInt = base16Scheme.CountColors()
	expectInt = 16
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	if base16Scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn() to be false got=%t", base16Scheme.ExtendedModeOn())
	}
}

func TestLoadBase16SchemeExtended(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark-extended.yaml")

	base16Scheme, err := base16.LoadBase16Scheme(testFile, ioutil.ReadFile)

	if err != nil {
		t.Errorf("expected no error err=%v", err)
	}

	gotString = base16Scheme.Author()
	expectString = "Chris Kempson (http://chriskempson.com)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = base16Scheme.Scheme()
	expectString = "Default Dark (Extended)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotInt = base16Scheme.CountColors()
	expectInt = 20
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	if !base16Scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn() to be true got=%t", base16Scheme.ExtendedModeOn())
	}
}

func TestLoadBase16SchemeError(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark-extended-invalid.yaml")

	_, err := base16.LoadBase16Scheme(testFile, ioutil.ReadFile)

	if err == nil {
		t.Errorf("expected error not nil")
	}
}

func TestGetColor(t *testing.T) {
	var gotColor base16.Color
	var expectColor base16.Color

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark.yaml")

	base16Scheme, _ := base16.LoadBase16Scheme(testFile, ioutil.ReadFile)

	gotColor = base16Scheme.GetColor("base0a")
	expectColor = 16239240
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
	gotColor = base16Scheme.GetColor("base0a")
	expectColor = 16239240
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}
