package knowledge_graph

import (
	"testing"

	"os"
	"encoding/json"
	"reflect"
)

var data = html()

func BenchmarkJSON(b *testing.B) {
	if data == nil { b.Fatal("No data!") }

  for i := 0; i < b.N; i++ {
  	g := New(data)
  	g.JSON()
  }
}

func TestJSON (t *testing.T) {
	expectedJSON := expectedResult()

	g := New(data)
	productedJSON := g.JSON()
	var parsedJSON map[string]interface{}

	err := json.Unmarshal(productedJSON, &parsedJSON)
	if err != nil { t.Errorf("%v", err) }

	if eq := reflect.DeepEqual(parsedJSON, expectedJSON); !eq {
		t.Error("Result is not equal to expected!")
	}
}

func html() []byte {
	data, err := os.ReadFile("../../files/van-gogh-paintings.html")

	if err != nil { panic(err) }

	return data
}

func expectedResult() map[string]interface{} {
	data, err := os.ReadFile("../../files/van-gogh-paintings-expected-array.json")
	if err != nil { panic(err) }

	var res = make(map[string]interface{})
	err = json.Unmarshal(data, &res)
	if err != nil { panic(err) }

	return res
}