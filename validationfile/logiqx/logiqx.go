/*
Package logiqx describes and handles the Logiqx XML validation file.
*/
package logiqx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/thalesvb/RPGL"
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

/*
ParseLogiqxXMLFile parses a XML file written with Logiqx schema and returns a
ValidationFile which can be queried to fetch information to build a playlist.
*/
func ParseLogiqxXMLFile(path string) RPGL.ValidationFile {
	var err error
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer xmlFile.Close()
	xmlFileByte, _ := ioutil.ReadAll(xmlFile)

	var dataFile logiqxDataFile
	err = xml.Unmarshal(xmlFileByte, &dataFile)
	if err != nil {
		println("Problem parsing validation file ", path)
		panic(err)
	}

	return dataFile
}
