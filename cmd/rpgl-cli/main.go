package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thalesvb/RPGL/ra"
	"github.com/thalesvb/RPGL/romfile"
	"github.com/thalesvb/RPGL/validationfile/mame"
)

func main() {

	flagCommon := flag.NewFlagSet("Common", flag.ContinueOnError)

	romDir := flagCommon.String("RomDir", "", "Root of ROM directory to be analysed")
	valFile := flagCommon.String("ValFile", "", "Validation file")
	playlistName := flagCommon.String("PlaylistName", "", "")

	var extensions = []string{".7z", ".zip"}

	flagCommon.Parse(os.Args[1:7])
	if *romDir == "" ||
		*valFile == "" ||
		*playlistName == "" {
		flagCommon.PrintDefaults()
		return
	}

	flags := ra.ParseFlags(os.Args[7:])
	validationFile := mame.ParseMameDatFile(*valFile)
	romFiles := romfile.FindRomsFromFolder(*romDir, extensions)
	fmt.Printf("Found %d file(s) that match ROM pattern\n", len(romFiles))
	fmt.Printf("Validation file contains %d game(s)\n", validationFile.Size())

	playlist := ra.BuildPlaylist(
		*playlistName,
		validationFile,
		romFiles,
		flags,
	)

	f, err := os.Create(*playlistName)
	if err != nil {
		panic(err)
	}
	f.Write(playlist.SerializePlaylist())
	f.Sync()
	f.Close()
}
