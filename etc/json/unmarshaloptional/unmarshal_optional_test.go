package unmarshaloptional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []byte(`
{
	"a": true,
	"b": "hi?",
	"d": 32
}
`)

type JsonStruct struct {
	A bool   `json:"a"`
	B string `json:"b"`
	C *bool  `json:"c,omitempty"`
	D int    `json:"d"`
}

func TestUnmarshalOptional(t *testing.T) {

	var dst JsonStruct
	err := json.Unmarshal(data, &dst)

	if assert.NoError(t, err) {
		assert.Nil(t, dst.C)
	}
}
