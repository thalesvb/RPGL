package mol_test

import (
	"strings"
	"testing"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/mol"
)

const emptyFile string = `
<?xml version="1.0" encoding="UTF-8"?>

<MyOwnLibrary>
    <header collection="COL" />
    <archives>
    </archives>
</MyOwnLibrary>
`
const mockFile string = `
<?xml version="1.0" encoding="UTF-8"?>

<MyOwnLibrary>
    <header collection="COL" />
    <archives>
        <archive file="GAME_FILE.EXT">
            <title>GAME_TITLE</title>
        </archive>
    </archives>
</MyOwnLibrary>
`
const (
	mckGameFile  string = "GAME_FILE.EXT"
	mckGameTitle string = "GAME_TITLE"
)

func loadFile(data string) RPGL.ValidationFile {
	mFile := strings.NewReader(data)
	file := mol.Parse(mFile)
	return file
}

func TestMinimalSuccess(t *testing.T) {
	file := loadFile(mockFile)
	if file.Size() == 0 {
		t.Errorf("XML not loaded")
	}
	meta := file.GetGameMetadata(mckGameFile)
	if meta.GetName() != mckGameTitle {
		t.Errorf("Oops!")
	}
	if meta.GetDescription() != mckGameTitle {
		t.Errorf("Oops!")
	}
}

func TestEmptyData(t *testing.T) {
	file := loadFile(emptyFile)
	if file.Size() != 0 {
		t.Errorf("Parser created an entry for a empty dataset")
	}
}
