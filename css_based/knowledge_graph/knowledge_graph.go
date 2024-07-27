package knowledge_graph

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"

	"serpapi-code-challenge-go/css_based/carousel"
)

type KnowledgeGraph struct {
	html []byte
}

func New(html []byte) KnowledgeGraph {
	return KnowledgeGraph{html: html}
}

func (g *KnowledgeGraph) JSON() []byte {
	carousel := carousel.New(g.htmlTree())

	res := []byte{}
	res = append(res, []byte(`{"artworks":`)...)
	res = append(res, carousel.JSON()...)
	res = append(res, []byte(`}`)...)

	return res
}

func (g *KnowledgeGraph) htmlTree() *goquery.Document {
	var reader bytes.Buffer
	reader.Write(g.html)

	doc, _ := goquery.NewDocumentFromReader(&reader)

	return doc
}