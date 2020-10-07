package gobase16

import (
	// "fmt"
	// "os"
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

	// fileKeys maps lower case key names (memory) to the case of the original key
	// names in the yaml file so that the original case is preserved when writing
	// files.
	fileKeys Base16FileKeys

	// // sortedColorNames contains all color names sorted alphabetically.
	// sortedColorNames []string

	// extendedMode is a flag which will be set when more than 16 colors are
	// defined.
	extendedMode bool
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
