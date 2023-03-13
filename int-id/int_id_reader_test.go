package intid

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

func TestIntIdReaderRead(t *testing.T) {
	id := 123

	paramProvider := new(ParamProviderMock)
	paramProvider.On("Param", "id").Return("123")

	sut := NewIntIdReader()

	result := sut.Read(paramProvider)

	assert.Equal(t, id, result)
}
