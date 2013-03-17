package xmlutils

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestGetXmlBody(t *testing.T) {
	xmlHeader := xml.Header //<?xml version="1.0" encoding="UTF-8"?>
	xmlBody := `<XMLBodySection>
	<AnElement>1</AnElement>
</XMLBodySection>
`
	xmlDoc := []byte(xmlHeader + xmlBody)

	testBody := StripHeader(xmlDoc)

	if xmlBytes := []byte(xmlBody); bytes.Compare(xmlBytes, testBody) != 0 {
		t.Errorf("XMLHeader not removed correcty.\nResult:\n%s\n\nExpected:\n%s", testBody, xmlBytes)
	}

}
