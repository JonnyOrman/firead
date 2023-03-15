package firead

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStructureDataDecoderDecode(t *testing.T) {
	data := make(map[string]interface{})
	data["prop1"] = "abc"
	data["prop2"] = 123

	sut := NewMapStructureDataDecoder[TestDocument]()

	result := sut.Decode(data)

	assert.Equal(t, "abc", result.Prop1)
	assert.Equal(t, 123, result.Prop2)
}
