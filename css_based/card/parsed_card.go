package card

// parsedCard is inner structure that used for collect
// parsed data and produce JSON
// independent of parsing order
// it can be fully private only if it placed in the 'card' package
type parsedCard struct {
	name []byte
	link []byte
	image []byte
	extension []byte
}

// TODO: try to create your custom fast JSON that will look more pretty?
func (c *parsedCard) JSON() []byte {
	res := []byte{}

	res = append(res, []byte(`{"name":"`)...)
	res = append(res, c.name...)
	res = append(res, []byte(`", "link":"`)...)
	res = append(res, c.link...)

	if len(c.image) == 0 {
		res = append(res, []byte(`", "image":`)...)
		res = append(res, []byte("null")...)
	} else {
		res = append(res, []byte(`", "image":"`)...)
		res = append(res, c.image...)
		res = append(res, []byte(`"`)...)
	}

	if len(c.extension) != 0 {
		res = append(res, []byte(`, "extensions": ["`)...)
		res = append(res, c.extension...)
		res = append(res, []byte(`"]`)...)
	}

	res = append(res, []byte(`}`)...)

	return res
}