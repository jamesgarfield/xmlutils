package xmlutils

import (
	"encoding/xml"
)

func StripHeader(xmlDoc []byte) []byte {
	lengthHeader := len(xml.Header)
	lengthDocument := len(xmlDoc)

	xmlBody := xmlDoc[lengthHeader:lengthDocument]
	return xmlBody
}
