package RPGL

import "path"

type GameMetadata interface {
	GetName() string
	GetDescription() string
}

type Playlist interface {
	GetName() string
	GetEntries() []PlaylistEntry
	SerializePlaylist() string
	AddEntry(PlaylistEntry)
}

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

func (r RomFile) FullPath() string {
	return path.Join(r.Path, (r.Name + r.Extension))
}

type ValidationFile interface {
	GetGameMetadata(romName string) GameMetadata
	Size() int
}
