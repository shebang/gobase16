package gobase16

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"sort"
	"strings"
	// "unicode"
	// "unicode/utf8"
)

// Base16Yaml declares the data structure for reading and writing a scheme file.
type Base16Yaml struct {
	Data       map[string]string `yaml:",omitempty,inline"`
	ColorNames []string
}

// UnmarshalBase16Yaml parses data as yaml and returns a Base16Yaml object on
// success.
func UnmarshalBase16Yaml(data []byte) (*Base16Yaml, error) {

	base16Yaml := Base16Yaml{Data: make(map[string]string)}

	err := yaml.Unmarshal(data, &base16Yaml)
	if err != nil {
		return nil, err
	}
	base16Yaml.ColorNames = base16Yaml.getYamlColorNames()
	if len(base16Yaml.ColorNames) > ExtendedModeMaxColors {
		return nil, fmt.Errorf("cannot use more than %d colors. Got %d colors", ExtendedModeMaxColors, len(base16Yaml.ColorNames))
	}

	return &base16Yaml, nil
}
func (y *Base16Yaml) getYamlColorNames() []string {

	colorNames := make([]string, 0, ExtendedModeMaxColors)
	for k := range y.Data {
		if strings.HasPrefix(k, "base") {
			colorNames = append(colorNames, k)
		}
	}
	sort.Strings(colorNames)
	return colorNames
}

// func TranslateBytesToLower(data []byte) {
// 	b := data
// 	for len(b) > 0 {
// 		r, size := utf8.DecodeRune(b)
// 		utf8.EncodeRune(b, unicode.ToLower(r))
// 		b = b[size:]
// 	}
// }
