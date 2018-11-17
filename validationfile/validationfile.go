package validationfile

import (
	"io"
	"os"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/mame"
)

/*
VFKind is a enumeration for structured files which contains validation data
*/
type VFKind string

const (
	kindMame VFKind = "MAME"
)

/*
Parse read file from system and returns a ValidationFile struct.
*/
func Parse(path string, vKind VFKind) RPGL.ValidationFile {
	file, err := os.Open(path)
	if err != nil {
		//logger.Error.Print("Error opening file:", err)
		panic(err)
	}
	defer file.Close()
	return parse(file, vKind)
}
func parse(file io.Reader, vKind VFKind) RPGL.ValidationFile {
	switch vKind {
	case kindMame:
		return mame.Parse(file)
	default:
		panic(vKind)
	}
}
