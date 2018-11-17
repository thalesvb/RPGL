package logiqx_test

import (
	"strings"
	"testing"

	"github.com/thalesvb/RPGL"
	"github.com/thalesvb/RPGL/validationfile/logiqx"
)

const emptyFile string = `
<?xml version="1.0"?>
<datafile />
`
const mockFile string = `
<?xml version="1.0"?>
<!DOCTYPE datafile PUBLIC "-//Logiqx//DTD ROM Management Datafile//EN" "http://www.logiqx.com/Dats/datafile.dtd">
<datafile>
	<header>
		<name>Logiqx</name>
		<description>Test File for Logiqx Parser</description>
		<category>Unit Test</category>
		<version>0.01</version>
	</header>
	<game name="MOCK">
		<description>MOCK_DESC</description>
	</game>
</datafile>
`

const (
	mckGameFile string = "MOCK"
	mckGameDesc string = "MOCK_DESC"
)

func loadFile(data string) RPGL.ValidationFile {
	return logiqx.Parse(strings.NewReader(data))
}

func TestEntryMissing(t *testing.T) {
	file := loadFile(emptyFile)
	entry := file.GetGameMetadata(mckGameFile)
	if entry != nil {
		t.Errorf("Oops")
	}
}

func TestParse(t *testing.T) {

	file := loadFile(mockFile)
	if file.Size() != 1 {
		t.Errorf("XML not parsed correctly")
	}
	meta := file.GetGameMetadata(mckGameFile)
	if meta.GetName() != mckGameFile {
		t.Errorf("Oops!")
	}
	if meta.GetDescription() != mckGameDesc {
		t.Errorf("")
	}
}
