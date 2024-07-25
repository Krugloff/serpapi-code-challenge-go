package thumbnails

import (
	"regexp"
	"bytes"
)

const elementPattern = `\(function\(\){var s='(?<image>[^']+?)';var ii=\['(?<id>[^']+?)'\];`
var elementRegexp = regexp.MustCompile(elementPattern)

type Thumbnails struct {
	html []byte

	value map[string][]byte
}

func New(html []byte) Thumbnails {
	return Thumbnails{html: html}
}

func (t *Thumbnails) FindBlob(id string) []byte {
	if t.value == nil { t.initValue() }

	return t.value[id]
}

func (t *Thumbnails) initValue() {
	t.value = make(map[string][]byte)

	els := elementRegexp.FindAllSubmatch(t.html, -1)

	for _, el := range els {
		blob := bytes.ReplaceAll(el[1], []byte(`\x`), []byte("x"))
		t.value[string(el[2])] = blob
	}
}