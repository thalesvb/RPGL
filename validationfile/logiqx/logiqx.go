/*
Package logiqx describes and handles the Logiqx XML validation file.
*/
package logiqx

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/logger"
)

/*
Representation of Logiqx XML data file.
*/
type logiqxDataFile struct {
	Header logiqxDataFileHeader `xml:"header"`
	Games  []logiqxGame         `xml:"game"`
}

func (df logiqxDataFile) GetGameMetadata(
	name string,
) RPGL.GameMetadata {
	for _, game := range df.Games {
		if game.Name == name {
			return game
		}
	}
	getLogger().Warning.Printf(`No metadata found for game file "%s".`, name)
	return nil
}
func (df logiqxDataFile) Size() int {
	return len(df.Games)
}

/*
Representation of the header tag.
*/
type logiqxDataFileHeader struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
}

/*
Representation of a rom tag.
*/
type logiqxRom struct {
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
	CRC  string `xml:"crc,attr"`
	SHA1 string `xml:"sha1,attr"`
}

/*
Representation of a game tag
*/
type logiqxGame struct {
	Name        string      `xml:"name,attr"`
	Description string      `xml:"description"`
	CloneOf     string      `xml:"cloneof,attr"`
	RomOf       string      `xml:"romof,attr"`
	Roms        []logiqxRom `xml:"rom"`
}

func (g logiqxGame) GetName() string {
	return g.Name
}
func (g logiqxGame) GetDescription() string {
	return g.Description
}

func getLogger() logger.Logger {
	return logger.GetLogger("validationfile.logiqx")
}

/*
ParseLogiqxXMLFile parses a XML file written with Logiqx schema and returns a
ValidationFile which can be queried to fetch information to build a playlist.
*/
func ParseLogiqxXMLFile(file io.Reader) RPGL.ValidationFile {
	var err error
	logger := getLogger()
	xmlFileByte, _ := ioutil.ReadAll(file)

	var dataFile logiqxDataFile
	err = xml.Unmarshal(xmlFileByte, &dataFile)
	if err != nil {
		logger.Error.Print("Problem parsing validation file")
		panic(err)
	}

	logger.Info.Printf("Validation file contains %d game(s)", len(dataFile.Games))
	return dataFile
}
