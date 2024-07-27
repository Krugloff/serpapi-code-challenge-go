package thumbnails

import (
	"bytes"
)

const (
	openBlobPattern = `(function(){var s='`
	openIDpattern = `';var ii=['`
	closeIDPattern = `'];`
)

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

	id, blob, chunk := t.nextPair(t.html);

	for len(chunk) != 0 {
		t.value[string(id)] = blob
		id, blob, chunk = t.nextPair(chunk);
	}
}

func (t *Thumbnails) nextPair(html []byte) (id []byte, blob []byte, rest []byte) {
	_, after, _ := bytes.Cut(html, []byte(openBlobPattern))
	blob, after, _ = bytes.Cut(after, []byte(openIDpattern))
	id, rest, _ = bytes.Cut(after, []byte(closeIDPattern))

	blob = bytes.ReplaceAll(blob, []byte(`\x`), []byte("x"))

	return
}


