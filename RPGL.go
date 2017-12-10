package RPGL

import "path/filepath"

/*
GameMetadata describes the minimal information required from a validation file to build a playlist.
*/
type GameMetadata interface {
	GetName() string
	GetDescription() string
}

/*
A Playlist is implemented by any concrete representation of a playlist to a specific software.
*/
type Playlist interface {
	GetName() string
	GetEntries() []PlaylistEntry
	SerializePlaylist() []byte
	AddEntry(PlaylistEntry)
}

/*
PlaylistFlags is implemented to fetch additional data to build a Playlist to as specific software.
*/
type PlaylistFlags interface {
	ParseFlags(args []string)
}

/*
PlaylistEntry is implemented by any concrete representation of a playlist's entry to a specific software.
*/
type PlaylistEntry interface {
	SerializeEntry() string
}

/*
RomFile is a basic description of a ROM file at OS level.
*/
type RomFile struct {
	Name      string
	Extension string
	Path      string
}

func (r RomFile) string() string {
	return r.Name + "|" + r.Extension + "|" + r.Path
}

/*
FullPath returns the absolute path of this ROM.
*/
func (r RomFile) FullPath() string {
	return filepath.Join(r.Path, (r.Name + r.Extension))
}

/*
A ValidationFile handles the operations in a Validation File to acquire some informations to build a Playlist.
*/
type ValidationFile interface {
	GetGameMetadata(romName string) GameMetadata
	Size() int
}
