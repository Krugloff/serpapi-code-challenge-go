package thumbnails

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

const (
	scriptSelector = `script:contains("function _setImagesSrc")`

	openBlobPattern = `(function(){var s='`
	openIDpattern = `';var ii=['`
	closeIDPattern = `'];`
)

type Thumbnails struct {
	doc *goquery.Document

	value map[string][]byte
}

func New(doc *goquery.Document) Thumbnails {
	return Thumbnails{doc: doc}
}

func (t *Thumbnails) FindBlob(id string) []byte {
	if t.value == nil { t.initValue() }

	return t.value[id]
}

func (t *Thumbnails) script() []byte {
	res := t.doc.FindMatcher(goquery.Single(scriptSelector)).Text()

	return []byte(res)
}

func (t *Thumbnails) initValue() {
	t.value = make(map[string][]byte)

	id, blob, chunk := t.nextPair(t.script());

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


