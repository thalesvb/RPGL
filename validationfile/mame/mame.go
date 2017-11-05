/*
Package mame handles MAME's .DAT validation file parsing.
*/
package mame

import (
	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/logiqx"
)

func ParseMameDatFile(path string) RPGL.ValidationFile {
	return logiqx.ParseLogiqxXmlFile(path)
}
