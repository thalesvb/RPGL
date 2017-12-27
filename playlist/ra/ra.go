/*
Package ra generates a Playlist for RetroArch system.
*/
package ra

import (
	"bytes"
	"flag"
	"strings"

	"github.com/thalesvb/RPGL"
)

/*
flagsRA represents a PlaylistFlags for RetroArch.
*/
type flagsRA struct {
	corePath string
	coreName string
}

func (f *flagsRA) ParseFlags(args []string) {
	fs := flag.NewFlagSet("RetroArch Playlist", flag.ContinueOnError)
	fs.StringVar(&f.coreName, "CoreName", "", "LibRetro Core name")
	fs.StringVar(&f.corePath, "CorePath", "", "LibRetro Core Path")
	if err := fs.Parse(args); err != nil {
		panic(err)
	}
}

/*
entryRA represents a PlaylistEntry of RetroArch.
*/
type entryRA struct {
	romPath  string
	romName  string
	corePath string
	coreName string
	crc      string
}

/*
SerializeEntry serializes a single entry following the RetroArch's Playlist file format.
*/
func (e entryRA) SerializeEntry() string {
	var crcLine string
	crcLine = e.crc + "|crc"
	aReturn := []string{
		e.romPath,
		e.romName,
		e.corePath,
		e.coreName,
		crcLine,
	}
	return strings.Join(aReturn, "\n")
}

/*
playlistRA represents a Playlist of RetroArch.
*/
type playlistRA struct {
	name    string
	entries []RPGL.PlaylistEntry
}

/*
AddEntry adds a game entry into playlist.
*/
func (p *playlistRA) AddEntry(e RPGL.PlaylistEntry) {
	p.entries = append(p.entries, e)
}

func (p playlistRA) GetEntries() []RPGL.PlaylistEntry {
	return p.entries
}

func (p playlistRA) GetName() string {
	return p.name
}

/*
SerializePlaylist serializes the entire playlist following the RetroArch's Playlist file format.
*/
func (p playlistRA) SerializePlaylist() []byte {

	var buffer bytes.Buffer

	for _, entry := range p.entries {
		buffer.WriteString(entry.SerializeEntry())
		buffer.WriteString("\n" + p.name + "\n")
	}
	buffer.WriteString("\n")
	return buffer.Bytes()
}

/*
buildPlaylist generates a roms' playlist for RA.
*/
func buildPlaylist(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	flags flagsRA,
) RPGL.Playlist {
	var romEntry entryRA
	var playlist = playlistRA{
		name: playlistName,
	}
	for _, rom := range roms {

		metadata := validationFile.GetGameMetadata(rom.Name)
		if metadata == nil {
			continue
		}

		romEntry = entryRA{
			romPath:  rom.FullPath(),
			romName:  metadata.GetDescription(),
			corePath: flags.corePath,
			coreName: flags.coreName,
			crc:      "0",
		}
		playlist.AddEntry(romEntry)
	}

	return &playlist
}

/*
ParseFlags parses RA playlist's specific flags.
*/
func parseFlags(args []string) flagsRA {
	flags := flagsRA{}
	flags.ParseFlags(args)
	return flags
}

/*
Build builds a RA playlist.
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
