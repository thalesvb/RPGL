/*
Package mame handles MAME's .DAT validation file parsing.
*/
package mame

import (
	"io"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/logiqx"
)

func ParseMameDatFile(file io.Reader) RPGL.ValidationFile {
	return logiqx.ParseLogiqxXMLFile(file)
}
