package card

import (
	"regexp"
	"bytes"
)

const (
	namePattern = `<a[^>]*?aria-label="(?<name>.+?)"`
	hrefPattern = `<a[^>]*?href="(?<link>.+?)"`
	titlePattern = `<a[^>]*?title="(?<title>.+?)"`
	thumbnailIDPattern = `<img[^>]*?id="(?<id>.+?)"`
)

var (
	nameRegexp = regexp.MustCompile(namePattern)
	hrefRegexp = regexp.MustCompile(hrefPattern)
	titleRegexp = regexp.MustCompile(titlePattern)
	thumbnailIDRegexp = regexp.MustCompile(thumbnailIDPattern)

	domain = []byte("https://www.google.com")
)

type BlobFinder interface {
	FindBlob(id string) []byte
}

type Card struct {
	html []byte
	thumbnails BlobFinder
}

func New(html []byte, thumbnails BlobFinder) Card {
	return Card{html: html, thumbnails: thumbnails}
}

func (c *Card) JSON() []byte {
	parsed := c.parsedCard()
	return parsed.JSON()
}

func (c *Card) parsedCard() parsedCard {
	return parsedCard{
		name: c.getName(),
		link: c.getLink(),
		image: c.getImage(),
		extension: c.getExtension()}
}

// TODO: potential issue?
// panic: runtime error: index out of range [1] with length 0
func (c *Card) getName() []byte {
	return nameRegexp.FindSubmatch(c.html)[1]
}

// TODO: potential issue?
// panic: runtime error: index out of range [1] with length 0
func (c *Card) getLink() []byte {
	link := hrefRegexp.FindSubmatch(c.html)[1]
	link = bytes.ReplaceAll(link, []byte(`&amp;`), []byte(`&`))
	link = append(domain, link...)
	return link
}

// TODO: potential issue?
// panic: runtime error: index out of range [1] with length 0
func (c *Card) getTitle() []byte {
	return titleRegexp.FindSubmatch(c.html)[1]
}

func (c *Card) getImage() []byte {
	return c.thumbnails.FindBlob(string(c.getThumbnailID()))
}

// TODO: potential issue?
// panic: runtime error: index out of range [1] with length 0
func (c *Card) getThumbnailID() []byte {
	return thumbnailIDRegexp.FindSubmatch(c.html)[1]
}

func (c *Card) getExtension() []byte {
	ext := bytes.Replace(c.getTitle(), c.getName(), []byte(""), 1)
	ext = bytes.TrimSpace(ext)
	ext = bytes.Trim(ext, "()")

	return ext
}