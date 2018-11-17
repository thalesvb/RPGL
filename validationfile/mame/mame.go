/*
Package mame handles MAME's .DAT validation file parsing.
*/
package mame

import (
	"io"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/logiqx"
)

func Parse(file io.Reader) RPGL.ValidationFile {
	return logiqx.Parse(file)
}
