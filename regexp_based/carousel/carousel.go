package carousel

import (
	"regexp"

	"serpapi-code-challenge-go/regexp_based/card"
	"serpapi-code-challenge-go/regexp_based/thumbnails"
)

const (
	blockPattern = `(?s)<g-scrolling-carousel.*?>(.*?)</g-scrolling-carousel>`
	elementPattern = `<a[^>]*?class="klitem".*?>.*?<\/a>`
)

var (
	blockRegexp = regexp.MustCompile(blockPattern)
	elementRegexp = regexp.MustCompile(elementPattern)
)

type Carousel struct {
	html []byte
}

func New(html []byte) Carousel {
	return Carousel{html: html}
}

// TODO: need add the check that something has been found
func (c *Carousel) JSON() []byte {
	thumbnails := thumbnails.New(c.html)
	els := elementRegexp.FindAll(c.blockHtml(), -1) // -1 means return w/o limit

	res := []byte(`[`)

	for i, el := range els {
		c := card.New(el, &thumbnails)
		res = append(res, c.JSON()...)

		if i+1 == len(els) { continue } // don't need add comma after last element
		res = append(res, []byte(`,`)...)
	}

	res = append(res, []byte(`]`)...)

	return res
}

// TODO: potential issue?
// panic: runtime error: index out of range [1] with length 0
func (c *Carousel) blockHtml() []byte {
	return blockRegexp.FindSubmatch(c.html)[1]
}