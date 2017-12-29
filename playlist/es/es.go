/*
Package es generates a Playlist for EmulationStation launcher.
https://github.com/Aloshi/EmulationStation/blob/master/GAMELISTS.md
*/
package es

import (
	"bytes"
	"encoding/xml"
	"flag"

	"github.com/thalesvb/RPGL"
)

const (
	outputXMLHeader string = `<?xml version="1.0"?>`
)

type flagsES struct {
	imagePath string
}

func (f *flagsES) ParseFlags(args []string) {
	fs := flag.NewFlagSet("EmulationStation Playlist", flag.ContinueOnError)
	fs.StringVar(&f.imagePath, "ImagePath", "", "Images Path")
	if err := fs.Parse(args); err != nil {
		panic(err)
	}
}

type esGame struct {
	XMLName xml.Name `xml:"game"`
	Path    string   `xml:"path"`
	Name    string   `xml:"name"`
	Image   string   `xml:"image,omitempty"`
}

func (g esGame) SerializeEntry() string {
	return ""
}

type esGameList struct {
	XMLName xml.Name `xml:"gameList"`
	Games   []esGame `xml:"game"`
}

type esGameListFile struct {
	GameList esGameList `xml:"gameList"`
}

func (p *esGameListFile) AddEntry(e RPGL.PlaylistEntry) {
	esGame := e.(esGame)
	p.GameList.Games = append(p.GameList.Games, esGame)
}

func (p esGameListFile) GetEntries() []RPGL.PlaylistEntry {
	var a interface{} = p.GameList.Games
	return a.([]RPGL.PlaylistEntry)
}

func (p esGameListFile) GetName() string {
	return ""
}

func (p esGameListFile) SerializePlaylist() []byte {
	output, err := xml.MarshalIndent(p.GameList, "", "\t")
	if err != nil {
		panic(err)
	}
	result := [][]byte{[]byte(outputXMLHeader), []byte("\n"), output}
	return bytes.Join(result, []byte(nil))
}

func buildPlaylist(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	flags flagsES,
) RPGL.Playlist {
	var romEntry esGame
	var playlist = esGameListFile{}
	for _, rom := range roms {

		metadata := validationFile.GetGameMetadata(rom.Name)
		if metadata == nil {
			continue
		}

		romEntry = esGame{
			Path: rom.FullPath(),
			Name: metadata.GetDescription(),
		}
		playlist.AddEntry(romEntry)

	}

	return &playlist
}

func parseFlags(args []string) flagsES {
	flags := flagsES{}
	flags.ParseFlags(args)
	return flags
}

/*
Build builds a ES playlist
*/
func Build(
	name string,
	additionalArgs []string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
) RPGL.Playlist {
	flags := parseFlags(additionalArgs)

	playlist := buildPlaylist(
		name,
		validationFile,
		roms,
		flags,
	)

	return playlist
}
