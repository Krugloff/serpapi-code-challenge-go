package card

import (
	"bytes"
)

const (
	closePattern = `"`
	nameOpenPattern = `aria-label="`
	hrefOpenPattern = `href="`
	titleOpenPattern = `title="`
	imageOpenPattern = `<img`
	thumbnailIDOpenPattern = `id="`
)

var (
	domain = []byte("https://www.google.com")
)

type BlobFinder interface {
	FindBlob(id string) []byte
}

type Card struct {
	doc []byte
	thumbnails BlobFinder
}

func New(doc []byte, thumbnails BlobFinder) Card {
	return Card{doc: doc, thumbnails: thumbnails}
}

func (c *Card) JSON() []byte {
	parsed := c.parsedCard()
	return parsed.JSON()
}

// see parsed_card.go in the same package
func (c *Card) parsedCard() parsedCard {
	name := c.getName()

	// the order is important since these methods used sequential scanning!
	return parsedCard{
		name: name,
		link: c.getLink(),
		extension: c.getExtension(name),
		image: c.getImage()}
}

// There is a dangerous behaviour - I make initial doc less and less actually

func (c *Card) getName() []byte {
	_, doc, _ := bytes.Cut(c.doc, []byte(nameOpenPattern))
	name, doc, _ := bytes.Cut(doc, []byte(closePattern))
	c.doc = doc
	return name
}

func (c *Card) getLink() []byte {
	_, doc, _ := bytes.Cut(c.doc, []byte(hrefOpenPattern))
	link, doc, _ := bytes.Cut(doc, []byte(closePattern))

	// this is little better than try to escape all chars
	// with the bytes -> string -> bytes conversations
	// link = []byte(html.UnescapeString(string(link)))
	link = bytes.ReplaceAll(link, []byte(`&amp;`), []byte(`&`))
	link = append(domain, link...)

	c.doc = doc
	return link
}

func (c *Card) getTitle() []byte {
	_, doc, _ := bytes.Cut(c.doc, []byte(titleOpenPattern))
	title, doc, _ := bytes.Cut(doc, []byte(closePattern))
	c.doc = doc
	return title
}

func (c *Card) getImage() []byte {
	return c.thumbnails.FindBlob(string(c.getThumbnailID()))
}

func (c *Card) getThumbnailID() []byte {
	_, doc, _ := bytes.Cut(c.doc, []byte(imageOpenPattern))
	_, doc, _ = bytes.Cut(doc, []byte(thumbnailIDOpenPattern))
	id, doc, _ := bytes.Cut(doc, []byte(closePattern))

	c.doc = doc
	return id
}

// getTitle will move the scanning position!
func (c *Card) getExtension(name []byte) []byte {
	res := bytes.Replace(c.getTitle(), name, []byte(""), 1)
	res = bytes.TrimSpace(res)
	res = bytes.Trim(res, "()")

	return res
}