package base16

import (
	// "fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	// ExtendedModeMaxColors specifies how many colors can be used in base16
	// extended mode.
	ExtendedModeMaxColors = 32
)

// Scheme is the internal representation of a base16 colors scheme. All color
// names are converted to lower case characters in order to avoid confusion when
// accessing color names.

type Scheme struct {
	// auther contains the author of the scheme
	author string

	// scheme holds the scheme identifier (or name)
	scheme string

	// colors holds all base16 colors in a string map. When accessing colors
	// keys are automatically converted to lower case characters.
	colors map[string]Color

	// originalKeys maps lower case key names to the case of the original key
	// names so that the original case is preserved when writing files
	originalKeys map[string]string

	// extendedMode is a flag which will be set when more than 16 colors are
	// defined.
	extendedMode bool
}

type FileReaderFunc func(string) ([]byte, error)

func LoadBase16Scheme(fname string, reader ...interface{}) (*Scheme, error) {
	var err error
	var data []byte
	var base16Yaml *Base16Yaml

	if _, err = os.Stat(fname); os.IsNotExist(err) {
		return nil, err
	}

	data, err = ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	base16Yaml, err = UnmarshalBase16Yaml(data)
	if err != nil {
		return nil, err
	}

	base16Scheme, err := fromBase16Yaml(base16Yaml)
	if err != nil {
		return nil, err
	}

	return base16Scheme, nil

}

func fromBase16Yaml(base16Yaml *Base16Yaml) (*Scheme, error) {
	extendedMode := false
	if len(base16Yaml.ColorNames) > 16 {
		extendedMode = true
	}

	scheme := Scheme{
		author:       base16Yaml.Data["author"],
		scheme:       base16Yaml.Data["scheme"],
		colors:       make(map[string]Color, len(base16Yaml.ColorNames)),
		originalKeys: make(map[string]string, len(base16Yaml.Data)),
		extendedMode: extendedMode,
	}

	colorName := ""
	for k, v := range base16Yaml.Data {
		if ValidColorName(k, extendedMode) {
			colorName = strings.ToLower(k)
			scheme.originalKeys[colorName] = k
			scheme.colors[colorName] = NewColor(v)
		} else {
			scheme.originalKeys[strings.ToLower(k)] = v
		}
	}

	return &scheme, nil
}

// Author returns the author of the scheme
func (scheme *Scheme) Author() string {
	return scheme.author
}

// Scheme returns the scheme identifier (name)
func (scheme *Scheme) Scheme() string {
	return scheme.scheme
}

// CountColors returns the number of colors
func (scheme *Scheme) CountColors() int {
	return len(scheme.colors)
}

// GetColor returns the color specified by colorname
func (scheme *Scheme) GetColor(colorname string) Color {
	return scheme.colors[strings.ToLower(colorname)]
}

// SetColor sets the color c for color name
func (scheme *Scheme) SetColor(c Color, colorname string) {
}

// ExtendedModeOn returns the extended mode flag
func (scheme *Scheme) ExtendedModeOn() bool {
	return scheme.extendedMode
}
