package carousel

import (
	_ "bytes"

	"github.com/PuerkitoBio/goquery"

	"serpapi-code-challenge-go/css_based/card"
	"serpapi-code-challenge-go/css_based/thumbnails"
)

const (
	blockSelector = `g-scrolling-carousel`
	cardSelector = `a.klitem`
)

type Carousel struct {
	doc *goquery.Document
}

func New(doc *goquery.Document) Carousel {
	return Carousel{doc: doc}
}

// TODO: need add the check that something has been found
func (c *Carousel) JSON() []byte {
	thumbnails := thumbnails.New(c.doc)

	res := []byte{}
	res = append(res, []byte(`[`)...)

	links := c.links()

	links.Each(func(i int, s *goquery.Selection) {
		cc := card.New(s, &thumbnails)
		res = append(res, cc.JSON()...)

		if i+1 < links.Length() { // till last element
			res = append(res, []byte(`,`)...)
		}
	})

	res = append(res, []byte(`]`)...)

	return res
}

func (c *Carousel) links() *goquery.Selection {
	return c.blockTree().Find(cardSelector)
}

func (c *Carousel) blockTree() *goquery.Selection {
	return c.doc.FindMatcher(goquery.Single(blockSelector))
}