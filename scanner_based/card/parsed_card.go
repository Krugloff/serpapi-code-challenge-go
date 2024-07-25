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

// this version produce much less allocations but looks like a liitle bit slower

// import "bytes"

// func (c *parsedCard) JSON() []byte {
// 	var res bytes.Buffer

// 	res.Write([]byte(`{"name":"`))
// 	res.Write(c.name)
// 	res.Write([]byte(`", "link":"`))
// 	res.Write(c.link)

// 	res.Write([]byte(`", "image":"`))
// 	res.Write(c.image)

// 	res.Write([]byte(`", "extensions": ["`))
// 	res.Write(c.extension)
// 	res.Write([]byte(`"]`))

// 	res.Write([]byte(`}`))

// 	return res.Bytes()
// }



// import "encoding/json"

// std marshalling is slow
// since it's based on reflection
// also it requires to have public fields
// these names you need to configure manually

// func (c *parsedCard) JSON() []byte {
// 	res, _ := json.Marshal(c)
// 	return res
// }



// Text template is slow and also requires public fields

// import "text/template"
// import "bytes"

// const (
// 	jsonTemplate = `{"name": "{{.Name}}", "link": "{{.Link}}", "image": "{{.Image}}", "extensions": ["{{.Extension}}"]}`
// )

// var (
// 	readyTemplate, _ = template.New("card").Parse(jsonTemplate)
// )

// func (c *parsedCard) JSON() []byte {
// 	// buffer implements io.Writer so it can be used instead os.Stdout
// 	var res bytes.Buffer
// 	readyTemplate.Execute(&res, c)
// 	return res.Bytes()
// }


