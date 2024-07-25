package knowledge_graph

import "serpapi-code-challenge-go/scanner_based/carousel"

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

// Template version
// import "text/template"

// const (
// 	jsonTemplate = `{"artworks": {{.}}}`
// )

// var (
// 	readyTemplate, _ = template.New("KnowledgeGraph").Parse(jsonTemplate)
// )

// func (g *KnowledgeGraph) JSON() []byte {
//  carousel := carousel.New(g.html)
// 	var res bytes.Buffer
// 	readyTemplate.Execute(&res, string(carousel.JSON()))
// 	return res.Bytes()
// }


// buffer version
// import "bytes"

// func (g *KnowledgeGraph) JSON() []byte {
// 	carousel := carousel.New(g.html)

// 	var res bytes.Buffer
// 	res.Write([]byte(`{"artworks":`))
// 	res.Write(carousel.JSON())
// 	res.Write([]byte(`}`))

// 	return res.Bytes()
// }