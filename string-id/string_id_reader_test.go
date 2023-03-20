package stringid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ParamProviderMock struct {
	mock.Mock
}

func (this ParamProviderMock) Param(key string) string {
	args := this.Called(key)
	return args.Get(0).(string)
}

func TestStringIdReaderRead(t *testing.T) {
	id := "abc"

	paramProvider := new(ParamProviderMock)
	paramProvider.On("Param", "id").Return(id)

	sut := NewStringIdReader()

	result := sut.Read(paramProvider)

	assert.Equal(t, id, result)
}
