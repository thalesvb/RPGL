package mame

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/thalesvb/RPGL"
)

type MameDataFile struct {
	Header MameDataFileHeader `xml:"header"`
	Games  []Game             `xml:"game"`
}

func (df MameDataFile) GetGameMetadata(
	name string,
) RPGL.GameMetadata {
	for _, game := range df.Games {
		if game.Name == name {
			return game
		}
	}
	return nil
}
func (df MameDataFile) Size() int {
	return len(df.Games)
}

type MameDataFileHeader struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
}

type MameRom struct {
	Name string `xml:"name,attr"`
	Size int    `xml:"size,attr"`
	CRC  string `xml:"crc,attr"`
	SHA1 string `xml:"sha1,attr"`
}

type Game struct {
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description"`
	CloneOf     string    `xml:"cloneof,attr"`
	RomOf       string    `xml:"romof,attr"`
	Roms        []MameRom `xml:"rom"`
}

func (g Game) GetName() string {
	return g.Name
}
func (g Game) GetDescription() string {
	return g.Description
}

func ParseMameDatFile(path string) RPGL.ValidationFile {
	var err error
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer xmlFile.Close()
	xmlFileByte, _ := ioutil.ReadAll(xmlFile)

	var dataFile MameDataFile
	err = xml.Unmarshal(xmlFileByte, &dataFile)
	if err != nil {
		println("Problem parsing validation file ", path)
		panic(err)
	}

	return dataFile
}
