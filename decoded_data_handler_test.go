package firead

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodedDataHandlerHandle(t *testing.T) {
	decodedData := make(map[string]interface{})

	sut := NewDecodedDataHandler[map[string]interface{}]()

	code, body := sut.Handle(decodedData)

	assert.Equal(t, 200, code)
	assert.Equal(t, decodedData, body)
}
