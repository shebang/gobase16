// +build test

package base16

import (
	"github.com/shebang/base16"
	"reflect"
	"testing"
)

var base16TestData = `
scheme: "Default Dark"
author: "Chris Kempson (http://chriskempson.com)"
base00: "181818"
base01: "282828"
base02: "383838"
base03: "585858"
base04: "b8b8b8"
base05: "d8d8d8"
base06: "e8e8e8"
base07: "f8f8f8"
base08: "ab4642"
base09: "dc9656"
base0A: "f7ca88"
base0B: "a1b56c"
base0C: "86c1b9"
base0D: "7cafc2"
base0E: "ba8baf"
base0F: "a16946"
`

// func TestBase16YamlTranslateToLower(t *testing.T) {
// 	var expectString string
// 	var gotString string

// 	testBytes := make([]byte, len(base16TestData))
// 	copy(testBytes, base16TestData)
// 	base16.TranslateBytesToLower(testBytes)

// 	expectString = "base0a"
// 	gotString = string(testBytes)[244:250]
// 	if gotString != expectString {
// 		t.Errorf("expected value=%s got=%s", expectString, gotString)
// 	}
// }

func TestBase16YamlLoading(t *testing.T) {

	base16Yaml, _ := base16.UnmarshalBase16Yaml([]byte(base16TestData))
	var gotString string

	testCases := []struct {
		colorName  string
		colorValue string
	}{
		{"base00", "181818"},
		{"base0F", "a16946"},
	}
	for _, table := range testCases {
		gotString = base16Yaml.Data[table.colorName]
		if gotString != table.colorValue {
			t.Errorf(
				"expected key=%s value=%s, got value=%s",
				table.colorName,
				table.colorValue,
				gotString,
			)
		}
	}
}

func TestGetYamlColorNames(t *testing.T) {

	base16Yaml, _ := base16.UnmarshalBase16Yaml([]byte(base16TestData))
	got := []string(base16Yaml.GetYamlColorNames())
	expect := []string{
		"base00",
		"base01",
		"base02",
		"base03",
		"base04",
		"base05",
		"base06",
		"base07",
		"base08",
		"base09",
		"base0A",
		"base0B",
		"base0C",
		"base0D",
		"base0E",
		"base0F",
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected value=%v, got=%v", expect, got)
	}
}
