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

type DataDecoderMock[T any] struct {
	mock.Mock
}

func (this DataDecoderMock[T]) Decode(data map[string]interface{}) T {
	args := this.Called(data)
	return args.Get(0).(T)
}

type DataHandlerMock[T any] struct {
	mock.Mock
}

func (this DataHandlerMock[T]) Handle(data T) (int, any) {
	args := this.Called(data)
	return args.Get(0).(int), args.Get(1).(any)
}

func TestExistingDocumentDataHandlerHandle(t *testing.T) {
	expectedCode := 200

	expectedBody := TestDocument{}

	data := make(map[string]interface{})
	data["key"] = "val"

	dataDecoder := new(DataDecoderMock[TestDocument])
	dataDecoder.On("Decode", data).Return(expectedBody)

	decodedDataHandler := new(DataHandlerMock[TestDocument])
	decodedDataHandler.On("Handle", expectedBody).Return(expectedCode, expectedBody)

	sut := NewExistingDocumentDataHandler[TestDocument](
		dataDecoder,
		decodedDataHandler,
	)

	code, body := sut.Handle(data)

	assert.Equal(t, expectedCode, code)
	assert.Equal(t, expectedBody, body)
}
