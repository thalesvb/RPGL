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

/*
SerializePlaylist serializes the entire playlist following the RetroArch's Playlist file format.
*/
func (p playlistRA) SerializePlaylist() string {

	var buffer bytes.Buffer

	for _, entry := range p.entries {
		buffer.WriteString(entry.SerializeEntry())
		buffer.WriteString("\n" + p.name + "\n")
	}
	buffer.WriteString("\n")
	return buffer.String()
}

/*
BuildPlaylist generates a playlist file for RetroArch.
*/
func BuildPlaylist(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	flags RPGL.PlaylistFlags,
) string {
	typedFlags := flags.(*flagsRA)
	return buildPlaylistInternal(playlistName, validationFile, roms, *typedFlags)
}

/*
buildPlaylistInternal generates a roms' playlist file for RA.
*/
func buildPlaylistInternal(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	flags flagsRA,
) string {
	var romEntry entryRA
	var playlist = playlistRA{
		name: playlistName,
	}
	for _, rom := range roms {

		metadata := validationFile.GetGameMetadata(rom.Name)
		if metadata == nil {
			println(rom.Name)
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

	return playlist.SerializePlaylist()
}

/*
ParseFlags parses RA playlist's specific flags.
*/
func ParseFlags(args []string) RPGL.PlaylistFlags {
	flags := flagsRA{}
	flags.ParseFlags(args)
	return &flags
}
