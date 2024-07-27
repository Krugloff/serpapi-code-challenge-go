package card

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

const (
	nameAttrName = `aria-label`
	hrefAttrName = `href`
	titleAttrName = `title`
	imgSelector = `img`
	thumbnailIDAttrName = `id`
)

var (
	domain = []byte("https://www.google.com")
)

type BlobFinder interface {
	FindBlob(id string) []byte
}

type Card struct {
	node *goquery.Selection
	thumbnails BlobFinder
}

func New(node *goquery.Selection, thumbnails BlobFinder) Card {
	return Card{node: node, thumbnails: thumbnails}
}

func (c *Card) JSON() []byte {
	parsed := c.parsedCard()
	return parsed.JSON()
}

// see parsed_card.go in the same package
func (c *Card) parsedCard() parsedCard {
	return parsedCard{
		name: c.getName(),
		link: c.getLink(),
		image: c.getImage(),
		extension: c.getExtension()}
}

func (c *Card) getName() []byte {
	res, _ := c.node.Attr(nameAttrName)
	return []byte(res)
}

func (c *Card) getLink() []byte {
	link, _ := c.node.Attr(hrefAttrName)
	res := bytes.ReplaceAll([]byte(link), []byte(`&amp;`), []byte(`&`))
	res = append(domain, res...)
	return res
}

func (c *Card) getTitle() []byte {
	res, _ := c.node.Attr(titleAttrName)
	return []byte(res)
}

func (c *Card) getImage() []byte {
	return c.thumbnails.FindBlob(string(c.getThumbnailID()))
}

func (c *Card) getThumbnailID() []byte {
	sel := c.node.FindMatcher(goquery.Single(imgSelector))
	attr, _ := sel.Attr(thumbnailIDAttrName)
	return []byte(attr)
}

func (c *Card) getExtension() []byte {
	ext := bytes.Replace(c.getTitle(), c.getName(), []byte(""), 1)
	ext = bytes.TrimSpace(ext)
	ext = bytes.Trim(ext, "()")

	return ext
}