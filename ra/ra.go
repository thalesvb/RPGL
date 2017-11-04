package ra

import (
	"bytes"
	"strings"

	"github.com/thalesvb/RPGL"
)

/*
EntryRA represents a playlist entry of RetroArch.
*/
type EntryRA struct {
	romPath  string
	romName  string
	corePath string
	coreName string
	crc      string
}

func (e EntryRA) SerializeEntry() string {
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

type PlaylistRA struct {
	name    string
	entries []RPGL.PlaylistEntry
}

func (p *PlaylistRA) AddEntry(e RPGL.PlaylistEntry) {
	p.entries = append(p.entries, e)
}

func (p PlaylistRA) SerializePlaylist() string {

	var buffer bytes.Buffer

	for _, entry := range p.entries {
		buffer.WriteString(entry.SerializeEntry())
		buffer.WriteString("\n" + p.name + "\n")
	}
	buffer.WriteString("\n")
	return buffer.String()
}

func BuildPlaylist(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	coreName string,
	corePath string,
) string {
	return buildMamePlaylist(playlistName, validationFile, roms, coreName, corePath)
}

/*
BuildMamePlaylist generates a playlist file of MAME roms for RetroArch.
*/
func buildMamePlaylist(
	playlistName string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
	coreName string,
	corePath string,
) string {
	var romEntry EntryRA
	var playlist = PlaylistRA{
		name: playlistName,
	}
	for _, rom := range roms {

		metadata := validationFile.GetGameMetadata(rom.Name)
		if metadata == nil {
			println(rom.Name)
			continue
		}

		romEntry = EntryRA{
			romPath:  rom.FullPath(),
			romName:  metadata.GetDescription(),
			corePath: corePath,
			coreName: coreName,
			crc:      "0",
		}
		playlist.AddEntry(romEntry)
	}

	return playlist.SerializePlaylist()
}
