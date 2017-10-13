package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/thalesvb/RPGL/romfile"
)

func main() {

	romDir := flag.String("RomDir", "", "Root of ROM directory to be analyzed")
	romExts := flag.String("RomExts", "", "Extensions used by ROMs")
	flag.Parse()

	if *romDir == "" ||
		*romExts == "" {
		flag.PrintDefaults()
		return
	}

	romExtensions := strings.Split(*romExts, `;`)

	romFiles := romfile.FindRomsFromFolder(*romDir, romExtensions)
	fmt.Printf("Found %d file(s) that match ROM pattern\n", len(romFiles))

}
