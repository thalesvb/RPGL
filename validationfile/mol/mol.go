package mol

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/logger"
)

type molFile struct {
	Header   molHeader    `xml:"header"`
	Archives []molArchive `xml:"archives>archive"`
}

type molArchive struct {
	File  string `xml:"file,attr"`
	Title string `xml:"title"`
}

func (md molArchive) GetDescription() string {
	return md.Title
}
func (md molArchive) GetName() string {
	return md.Title
}

type molHeader struct {
	Collection string `xml:"collection,attr"`
}

func (mf molFile) GetGameMetadata(
	name string,
) RPGL.GameMetadata {
	for _, dump := range mf.Archives {
		if dump.File == name {
			return dump
		}
	}
	getLogger().Warning.Printf(`No metadata found for game file "%s".`, name)
	return nil
}
func (mf molFile) Size() int {
	return len(mf.Archives)
}

func getLogger() logger.Logger {
	return logger.GetLogger("validationfile.mol")
}

/*Parse receives a data reader of an MOL file and returns a ValidationFile struct.*/
func Parse(file io.Reader) RPGL.ValidationFile {
	var err error
	logger := getLogger()
	xmlFileByte, _ := ioutil.ReadAll(file)

	var mol molFile
	err = xml.Unmarshal(xmlFileByte, &mol)
	if err != nil {
		logger.Error.Print("Problem parsing validation file")
		panic(err)
	}

	logger.Info.Printf("Validation file contains %d game(s)", len(mol.Archives))
	return mol
}
