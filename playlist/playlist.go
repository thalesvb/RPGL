package playlist

import (
	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/playlist/ra"
)

const (
	systemRA string = "RetroArch"
)

/*
Build builds a playlist for a system.
*/
func Build(
	playlistSystem string,
	name string,
	additionalArgs []string,
	validationFile RPGL.ValidationFile,
	roms []RPGL.RomFile,
) RPGL.Playlist {

	switch playlistSystem {
	case systemRA:
		return ra.Build(
			name,
			additionalArgs,
			validationFile,
			roms,
		)
	default:
		panic(playlistSystem)
	}

}
