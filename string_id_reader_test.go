package firead

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringIdReaderRead(t *testing.T) {
	id := "abc"

	paramProvider := new(ParamProviderMock)
	paramProvider.On("Param", "id").Return(id)

	sut := NewStringIdReader()

	result := sut.Read(paramProvider)

	assert.Equal(t, id, result)
}
