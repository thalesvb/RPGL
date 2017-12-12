package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thalesvb/RPGL/playlist"
	"github.com/thalesvb/RPGL/romfile"
	"github.com/thalesvb/RPGL/validationfile/mame"
)

func main() {

	flagCommon := flag.NewFlagSet("Common", flag.ContinueOnError)

	romDir := flagCommon.String("RomDir", "", "Root of ROM directory to be analysed")
	valFile := flagCommon.String("ValFile", "", "Validation file")
	playlistName := flagCommon.String("PlaylistName", "", "Playlist's name")
	playlistSystem := flagCommon.String("PlaylistSystem", "", "System to generate the playlist")

	var extensions = []string{".7z", ".zip"}

	// 1 (offset) + 2 (flag name / value) * (quantity of mandatory common flags)
	const idxSplitCommonAdditionalArgs = 1 + 2*4

	commonArgs := os.Args[1:idxSplitCommonAdditionalArgs]
	additionalArgs := os.Args[idxSplitCommonAdditionalArgs:]

	flagCommon.Parse(commonArgs)
	if *romDir == "" ||
		*valFile == "" ||
		*playlistName == "" ||
		*playlistSystem == "" {
		flagCommon.PrintDefaults()
		return
	}

	validationFile := mame.ParseMameDatFile(*valFile)
	romFiles := romfile.FindRomsFromFolder(*romDir, extensions)
	fmt.Printf("Found %d file(s) that match ROM pattern\n", len(romFiles))
	fmt.Printf("Validation file contains %d game(s)\n", validationFile.Size())

	playlist := playlist.Build(
		*playlistSystem,
		*playlistName,
		additionalArgs,
		validationFile,
		romFiles,
	)

	f, err := os.Create(*playlistName)
	if err != nil {
		panic(err)
	}
	f.Write(playlist.SerializePlaylist())
	f.Sync()
	f.Close()
}
