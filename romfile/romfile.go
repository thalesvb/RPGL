package romfile

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/thalesvb/RPGL/logger"

	"github.com/thalesvb/RPGL"
)

/*
FindRomsFromFolder tries to recursively match possible ROM files based on extensions list.
It works like "find all files in this directory and it's subdirectiories with those extensions".
*/
func FindRomsFromFolder(rootPath string, extensions []string) []RPGL.RomFile {
	sort.Strings(extensions)

	for i := range extensions {
		var ext = &extensions[i]
		*ext = strings.ToUpper(*ext)
	}

	return findRomsFromRootFolderInternal(
		rootPath,
		extensions,
	)
}

/*
Internal implementation of FindRomsFromFolder function.
*/
func findRomsFromRootFolderInternal(rootPath string, extensions []string) []RPGL.RomFile {
	var aRomFile RPGL.RomFile
	var romFiles []RPGL.RomFile
	var baseName string
	var extsRightBound = len(extensions)

	dirEntries, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil
	}

	for _, fi := range dirEntries {
		baseName = fi.Name()

		if fi.IsDir() == true {
			romFiles = append(
				romFiles,
				findRomsFromRootFolderInternal(
					filepath.Join(rootPath, baseName), extensions,
				)...,
			)
			continue
		}

		fileExt := filepath.Ext(baseName)
		fileExtUpper := strings.ToUpper(fileExt)
		idx := sort.SearchStrings(extensions, fileExtUpper)
		if (idx < extsRightBound) && (extensions[idx] == fileExtUpper) {

			baseName = strings.Replace(baseName, fileExt, "", -1)

			aRomFile = RPGL.RomFile{
				Name:      baseName,
				Extension: fileExt,
				Path:      rootPath,
			}
			romFiles = append(romFiles, aRomFile)
		}
	}

	getLogger().Info.Printf("Found %d file(s) that match ROM pattern\n", len(romFiles))
	return romFiles

}

func getLogger() logger.Logger {
	return logger.GetLogger("romfile")
}
