package firead

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestDocument struct {
	Prop1 string
	Prop2 int
}

type DataRetrieverMock[TId Id] struct {
	mock.Mock
}

func (this DataRetrieverMock[TId]) Retrieve(id TId) map[string]interface{} {
	args := this.Called(id)
	return args.Get(0).(map[string]interface{})
}

func TestFirestoreDocumentReaderRead(t *testing.T) {
	id := "abc"

	data := make(map[string]interface{})
	data["prop1"] = "def"
	data["prop2"] = 123

	dataRetriever := new(DataRetrieverMock[string])
	dataRetriever.On("Retrieve", id).Return(data)

	sut := NewFirestoreDocumentReader[TestDocument, string](dataRetriever)

	expectedResult := TestDocument{}
	expectedResult.Prop1 = "def"
	expectedResult.Prop2 = 123

	result := sut.Read(id)

	assert.Equal(t, expectedResult, result)
}
