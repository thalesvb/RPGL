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

type LogiqxDataFile struct {
	Header LogiqxDataFileHeader `xml:"header"`
	Games  []LogiqxGame         `xml:"game"`
}

func (df LogiqxDataFile) GetGameMetadata(
	name string,
) RPGL.GameMetadata {
	for _, game := range df.Games {
		if game.Name == name {
			return game
		}
	}
	return nil
}
func (df LogiqxDataFile) Size() int {
	return len(df.Games)
}

type LogiqxDataFileHeader struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
}

type LogiqxRom struct {
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
	CRC  string `xml:"crc,attr"`
	SHA1 string `xml:"sha1,attr"`
}

type LogiqxGame struct {
	Name        string      `xml:"name,attr"`
	Description string      `xml:"description"`
	CloneOf     string      `xml:"cloneof,attr"`
	RomOf       string      `xml:"romof,attr"`
	Roms        []LogiqxRom `xml:"rom"`
}

func (g LogiqxGame) GetName() string {
	return g.Name
}
func (g LogiqxGame) GetDescription() string {
	return g.Description
}

func ParseLogiqxXmlFile(path string) RPGL.ValidationFile {
	var err error
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer xmlFile.Close()
	xmlFileByte, _ := ioutil.ReadAll(xmlFile)

	var dataFile LogiqxDataFile
	err = xml.Unmarshal(xmlFileByte, &dataFile)
	if err != nil {
		println("Problem parsing validation file ", path)
		panic(err)
	}

	return dataFile
}
