package carousel

import (
	"bytes"
	"serpapi-code-challenge-go/scanner_based/card"
	"serpapi-code-challenge-go/scanner_based/thumbnails"
)

const (
	blockOpenPattern = `<g-scrolling-carousel`
	blockClosePattern = `</g-scrolling-carousel`

	linkOpenPattern = `<a`
	linkClosePattern = `</a`
)

type Carousel struct {
	html []byte
}

func New(html []byte) Carousel {
	return Carousel{html: html}
}

// TODO: need add the check that something has been found
func (c *Carousel) JSON() []byte {
	_, after, _ := bytes.Cut(c.html, []byte(blockOpenPattern))
	blockHtml, after, _ := bytes.Cut(after, []byte(blockClosePattern))

	// so it will search script functions inside full page
	thumbnails := thumbnails.New(after)

	res := []byte{}
	res = append(res, []byte(`[`)...)

	// produce json for each found link element
	link, chunk := c.nextLink(blockHtml)

	for {
		cc := card.New(link, &thumbnails)
		res = append(res, cc.JSON()...)

		link, chunk = c.nextLink(chunk)

		if len(chunk) == 0 { break }

		res = append(res, []byte(`,`)...)
	}
	//

	res = append(res, []byte(`]`)...)

	return res
}

// TODO it will be good to check that this is a klitem element
// I can't add this check to openPattern because it will Cut everything before css class definition
func (c *Carousel) nextLink(html []byte) (link []byte, rest []byte) {
	_, after, _ := bytes.Cut(html, []byte(linkOpenPattern))
	link, rest, _ = bytes.Cut(after, []byte(linkClosePattern))

	return // naked return
}