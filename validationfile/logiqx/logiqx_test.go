package logiqx_test

import (
	"strings"
	"testing"

	"github.com/thalesvb/RPGL/validationfile/logiqx"
)

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

func TestLogiqxParse(t *testing.T) {

	file := logiqx.ParseLogiqxXMLFile(strings.NewReader(mockFile))
	if file.Size() != 1 {
		t.Errorf("XML not parsed correctly")
	}
	meta := file.GetGameMetadata("MOCK")
	if meta.GetName() != "MOCK" {
		t.Errorf("Oops!")
	}
	if meta.GetDescription() != "MOCK_DESC" {
		t.Errorf("")
	}
}
