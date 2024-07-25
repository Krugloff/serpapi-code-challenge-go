package knowledge_graph

import "serpapi-code-challenge-go/regexp_based/carousel"

type KnowledgeGraph struct {
	html []byte
}

func New(html []byte) KnowledgeGraph {
	return KnowledgeGraph{html: html}
}

func (g *KnowledgeGraph) JSON() []byte {
	carousel := carousel.New(g.html)

	res := []byte{}
	res = append(res, []byte(`{"artworks":`)...)
	res = append(res, carousel.JSON()...)
	res = append(res, []byte(`}`)...)

	return res
}