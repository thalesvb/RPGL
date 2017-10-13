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

	romDir := flag.String("RomDir", "", "Root of ROM directory to be analysed")
	valFile := flag.String("ValFile", "", "Validation file")
	coreName := flag.String("CoreName", "", "LibRetro Core name")
	corePath := flag.String("CorePath", "", "LibRetro Core Path")
	playlistName := flag.String("PlaylistName", "", "")

	var extensions = []string{".7z", ".zip"}

	flag.Parse()
	if *romDir == "" ||
		*valFile == "" ||
		*coreName == "" ||
		*corePath == "" ||
		*playlistName == "" {
		flag.PrintDefaults()
		return
	}

	validationFile := mame.ParseMameDatFile(*valFile)
	romFiles := romfile.FindRomsFromFolder(*romDir, extensions)
	fmt.Printf("Found %d file(s) that match ROM pattern\n", len(romFiles))
	fmt.Printf("Validation file contains %d game(s)\n", validationFile.Size())

	playlistData := ra.BuildPlaylist(
		*playlistName,
		validationFile,
		romFiles,
		*coreName,
		*corePath,
	)

	f, err := os.Create(*playlistName)
	if err != nil {
		panic(err)
	}
	f.WriteString(playlistData)
	f.Sync()
	f.Close()
}
